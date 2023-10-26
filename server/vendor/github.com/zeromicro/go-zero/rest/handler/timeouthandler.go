package handler

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/rest/internal"
)

const (
	statusClientClosedRequest = 499
	reason                    = "Request Timeout"
	headerUpgrade             = "Upgrade"
	valueWebsocket            = "websocket"
)

// TimeoutHandler returns the handler with given timeout.
// If client closed request, code 499 will be logged.
// Notice: even if canceled in server side, 499 will be logged as well.
func TimeoutHandler(duration time.Duration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		if duration > 0 {
			return &timeoutHandler{
				handler: next,
				dt:      duration,
			}
		}

		return next
	}
}

// timeoutHandler is the handler that controls the request timeout.
// Why we implement it on our own, because the stdlib implementation
// treats the ClientClosedRequest as http.StatusServiceUnavailable.
// And we write the codes in logs as code 499, which is defined by nginx.
type timeoutHandler struct {
	handler http.Handler
	dt      time.Duration
}

func (h *timeoutHandler) errorBody() string {
	return reason
}

func (h *timeoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get(headerUpgrade) == valueWebsocket {
		h.handler.ServeHTTP(w, r)
		return
	}

	ctx, cancelCtx := context.WithTimeout(r.Context(), h.dt)
	defer cancelCtx()

	r = r.WithContext(ctx)
	done := make(chan struct{})
	tw := &timeoutWriter{
		w:   w,
		h:   make(http.Header),
		req: r,
	}
	panicChan := make(chan interface{}, 1)
	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		h.handler.ServeHTTP(tw, r)
		close(done)
	}()
	select {
	case p := <-panicChan:
		panic(p)
	case <-done:
		tw.mu.Lock()
		defer tw.mu.Unlock()
		dst := w.Header()
		for k, vv := range tw.h {
			dst[k] = vv
		}
		if !tw.wroteHeader {
			tw.code = http.StatusOK
		}
		w.WriteHeader(tw.code)
		w.Write(tw.wbuf.Bytes())
	case <-ctx.Done():
		tw.mu.Lock()
		defer tw.mu.Unlock()
		// there isn't any user-defined middleware before TimoutHandler,
		// so we can guarantee that cancelation in biz related code won't come here.
		httpx.Error(w, ctx.Err(), func(w http.ResponseWriter, err error) {
			if errors.Is(err, context.Canceled) {
				w.WriteHeader(statusClientClosedRequest)
			} else {
				w.WriteHeader(http.StatusServiceUnavailable)
			}
			io.WriteString(w, h.errorBody())
		})
		tw.timedOut = true
	}
}

type timeoutWriter struct {
	w    http.ResponseWriter
	h    http.Header
	wbuf bytes.Buffer
	req  *http.Request

	mu          sync.Mutex
	timedOut    bool
	wroteHeader bool
	code        int
}

var _ http.Pusher = (*timeoutWriter)(nil)

// Header returns the underline temporary http.Header.
func (tw *timeoutWriter) Header() http.Header { return tw.h }

// Push implements the Pusher interface.
func (tw *timeoutWriter) Push(target string, opts *http.PushOptions) error {
	if pusher, ok := tw.w.(http.Pusher); ok {
		return pusher.Push(target, opts)
	}
	return http.ErrNotSupported
}

// Write writes the data to the connection as part of an HTTP reply.
// Timeout and multiple header written are guarded.
func (tw *timeoutWriter) Write(p []byte) (int, error) {
	tw.mu.Lock()
	defer tw.mu.Unlock()

	if tw.timedOut {
		return 0, http.ErrHandlerTimeout
	}

	if !tw.wroteHeader {
		tw.writeHeaderLocked(http.StatusOK)
	}
	return tw.wbuf.Write(p)
}

func (tw *timeoutWriter) writeHeaderLocked(code int) {
	checkWriteHeaderCode(code)

	switch {
	case tw.timedOut:
		return
	case tw.wroteHeader:
		if tw.req != nil {
			caller := relevantCaller()
			internal.Errorf(tw.req, "http: superfluous response.WriteHeader call from %s (%s:%d)",
				caller.Function, path.Base(caller.File), caller.Line)
		}
	default:
		tw.wroteHeader = true
		tw.code = code
	}
}

func (tw *timeoutWriter) WriteHeader(code int) {
	tw.mu.Lock()
	defer tw.mu.Unlock()
	tw.writeHeaderLocked(code)
}

func checkWriteHeaderCode(code int) {
	if code < 100 || code > 599 {
		panic(fmt.Sprintf("invalid WriteHeader code %v", code))
	}
}

// relevantCaller searches the call stack for the first function outside of net/http.
// The purpose of this function is to provide more helpful error messages.
func relevantCaller() runtime.Frame {
	pc := make([]uintptr, 16)
	n := runtime.Callers(1, pc)
	frames := runtime.CallersFrames(pc[:n])
	var frame runtime.Frame
	for {
		frame, more := frames.Next()
		if !strings.HasPrefix(frame.Function, "net/http.") {
			return frame
		}
		if !more {
			break
		}
	}
	return frame
}
