package errorx

type ErrorCode uint32

//全局错误码
const (
	SERVER_COMMON_ERROR           ErrorCode = 1001
	REUQEST_PARAM_ERROR           ErrorCode = 1002
	TOKEN_EXPIRE_ERROR            ErrorCode = 1003
	TOKEN_GENERATE_ERROR          ErrorCode = 1004
	DB_ERROR                      ErrorCode = 1005
	DB_UPDATE_AFFECTED_ZERO_ERROR ErrorCode = 1006
	USER_NOT_FOUND                ErrorCode = 1007
)

var message = map[ErrorCode]string{
	SERVER_COMMON_ERROR:           "服务内部错误",
	REUQEST_PARAM_ERROR:           "请求参数错误",
	TOKEN_EXPIRE_ERROR:            "TOKEN过期",
	TOKEN_GENERATE_ERROR:          "TOKEN生成失败",
	DB_ERROR:                      "数据库错误",
	DB_UPDATE_AFFECTED_ZERO_ERROR: "更新数据影响行数为0",
	USER_NOT_FOUND:                "用户户不存在",
}

func ErrorxMessage(errorCode ErrorCode) string {
	if msg, ok := message[errorCode]; ok {
		return msg
	} else {
		return "错误未定义，请查看日志"
	}
}

func IsErrorxCode(errcode ErrorCode) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
