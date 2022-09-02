package errorx

import (
	"fmt"
)

type Errorx struct {
	Code ErrorCode
	Msg  string
}

func (e *Errorx) Error() string {
	return fmt.Sprintf("ErrorxCode: %d, ErrorMsg: %s", e.Code, e.Msg)
}

func New(code ErrorCode, msg string) error {
	return &Errorx{
		Code: code,
		Msg:  msg,
	}
}

//func (e *Errorx) ErrorxCode() uint32 {
//	return uint32(e.code)
//}
//
//func (e *Errorx) ErrorxMsg() string {
//	return e.msg
//}

func NewWithCode(code ErrorCode) *Errorx {
	return &Errorx{
		Code: code,
		Msg:  ErrorxMessage(code),
	}
}

//
//func NewWithMsg(msg string) *Errorx {
//	return &Errorx{
//		code: SERVER_COMMON_ERROR,
//		msg:  msg,
//	}
//}
