package errorx

import (
	"encoding/json"
	"fmt"
)

type Errorx struct {
	Code ErrorCode
	Msg  string
}

func (e *Errorx) Error() string {
	if eb, err := json.Marshal(e); err != nil {
		return fmt.Sprintf("code: %d, message: %s", e.Code, e.Msg)
	} else {
		return string(eb)
	}

}

func New(code ErrorCode, msg string) error {
	return &Errorx{
		Code: code,
		Msg:  msg,
	}
}

func (e *Errorx) ErrorxCode() uint32 {
	return uint32(e.Code)
}

func (e *Errorx) ErrorxMsg() string {
	return e.Msg
}

func NewWithCode(code ErrorCode) *Errorx {
	return &Errorx{
		Code: code,
		Msg:  ErrorxMessage(code),
	}
}

func NewCommonCodeWithMsg(msg string) *Errorx {
	return &Errorx{
		Code: SERVER_COMMON_ERROR,
		Msg:  msg,
	}
}
