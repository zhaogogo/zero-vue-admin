package errorx

import (
	"fmt"
	"sync"
)

type Errorx struct {
	Code  ErrorCode
	Msg   string
	Meta  []interface{}
	Cause error
	mu    sync.Mutex
}

func (e *Errorx) Error() string {
	return fmt.Sprintf("x error: code = %v message: %s cause: %v", e.Code, e.Msg, e.Cause)
}

func (e *Errorx) WithMeta(rpcMethod string, rpcerr string, param interface{}) *Errorx {
	e.mu.Lock()
	defer e.mu.Unlock()
	if rpcMethod == "" && rpcerr == "" {
		e.Meta = append(e.Meta, param)
		return e
	}
	e.Meta = append(
		e.Meta,
		map[string]interface{}{
			"rpc方法": rpcMethod,
			"rpc错误": rpcerr, //如果rocerr类型是error json marshal
			"请求参数":  param,
		},
	)
	return e
}

func (e *Errorx) Message() string {
	return fmt.Sprintf("%s", e.Msg)
}

func New(causeErr error, msg string) *Errorx {
	return &Errorx{
		Code:  SERVER_COMMON_ERROR,
		Msg:   msg,
		Cause: causeErr,
	}
}

func NewByCode(causeErr error, code ErrorCode) *Errorx {
	return &Errorx{
		Code:  code,
		Msg:   ErrorxMessage(code),
		Cause: causeErr,
	}
}
