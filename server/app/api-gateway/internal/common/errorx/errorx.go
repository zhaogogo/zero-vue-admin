package errorx

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"sync"
)

type Errorx struct {
	Code  ErrorCode
	Msg   string
	Cause error
}

func (e *Errorx) Error() string {
	return fmt.Sprintf("x error: code = %v message: %s cause: %v", e.Code, e.Msg, e.Cause)
}

func New(causeErr error, code ErrorCode, msg string) error {
	return &Errorx{
		Code:  code,
		Msg:   msg,
		Cause: causeErr,
	}
}

func NewFromCode(causeErr error, code ErrorCode) error {
	return &Errorx{
		Code:  code,
		Msg:   ErrorxMsg(code),
		Cause: causeErr,
	}
}

type MsgErrList struct {
	List []string
	mu   sync.Mutex
}

func (m *MsgErrList) Append(v string) {
	logx.Error(v)
	m.mu.Lock()
	defer m.mu.Unlock()
	m.List = append(m.List, v)
}
