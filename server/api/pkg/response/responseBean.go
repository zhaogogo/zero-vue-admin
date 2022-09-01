package response

type ResponseBean struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func BeanOKWithData(data interface{}) *ResponseBean {
	return &ResponseBean{
		Code: 0,
		Msg:  "OK",
		Data: data,
	}
}

func BeanFail(code uint32, msg string) *ResponseBean {
	return &ResponseBean{
		Code: code,
		Msg:  msg,
	}
}

func BeanFailWithMessage(msg string) *ResponseBean {
	return &ResponseBean{
		Code: 500,
		Msg:  msg,
		Data: nil,
	}
}
