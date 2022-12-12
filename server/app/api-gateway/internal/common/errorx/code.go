package errorx

type ErrorCode uint32

//全局错误码
const (
	SERVER_COMMON_ERROR = ErrorCode(1001 + iota)
	REUQEST_PARAM_ERROR
	TOKEN_EXPIRE_ERROR
	TOKEN_GENERATE_ERROR
	DB_ERROR
	DB_NOTFOUND
	DB_UPDATE_AFFECTED_ZERO_ERROR
	USER_NOT_FOUND
	USER_PERMISSION_REJECT
)

var message = map[ErrorCode]string{
	SERVER_COMMON_ERROR:           "服务内部错误",
	REUQEST_PARAM_ERROR:           "请求参数错误",
	TOKEN_EXPIRE_ERROR:            "TOKEN过期",
	TOKEN_GENERATE_ERROR:          "TOKEN生成失败",
	DB_ERROR:                      "数据库错误",
	DB_NOTFOUND:                   "数据库无记录",
	DB_UPDATE_AFFECTED_ZERO_ERROR: "更新数据影响行数为0",
	USER_NOT_FOUND:                "用户户不存在",
	USER_PERMISSION_REJECT:        "用户权限被拒绝",
}

func ErrorxMsg(errorCode ErrorCode) string {
	if msg, ok := message[errorCode]; ok {
		return msg
	} else {
		return "错误未定义"
	}
}
