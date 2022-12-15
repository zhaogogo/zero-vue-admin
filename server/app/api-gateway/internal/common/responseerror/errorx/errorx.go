package errorx

import (
	"fmt"
)

type Errorx struct {
	Code  ErrorCode
	Msg   string
	Cause error
}

func (e *Errorx) Error() string {
	return fmt.Sprintf("x error: code = %v message: %s cause: %v", e.Code, e.Msg, e.Cause)
}

func (e *Errorx) Message() string {
	return fmt.Sprintf("%s", e.Msg)
}

func New(causeErr error, errorxcode ErrorCode, msg string) error {
	return &Errorx{
		Code:  errorxcode,
		Msg:   msg,
		Cause: causeErr,
	}
}

func NewByCode(causeErr error, code ErrorCode) error {
	return &Errorx{
		Code:  code,
		Msg:   ErrorxMessage(code),
		Cause: causeErr,
	}
}
