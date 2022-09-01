package errorx

type ErrorCode uint32

//全局错误码
const SERVER_COMMON_ERROR ErrorCode = 100001
const REUQEST_PARAM_ERROR ErrorCode = 100002
const TOKEN_EXPIRE_ERROR ErrorCode = 100003
const TOKEN_GENERATE_ERROR ErrorCode = 100004
const DB_ERROR ErrorCode = 100005
const DB_UPDATE_AFFECTED_ZERO_ERROR ErrorCode = 100006

var message = map[ErrorCode]string{
	SERVER_COMMON_ERROR:           "服务内部错误",
	REUQEST_PARAM_ERROR:           "请求参数错误",
	TOKEN_EXPIRE_ERROR:            "TOKEN过期",
	TOKEN_GENERATE_ERROR:          "TOKEN生成失败",
	DB_ERROR:                      "数据库错误",
	DB_UPDATE_AFFECTED_ZERO_ERROR: "更新数据影响行数为0",
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
