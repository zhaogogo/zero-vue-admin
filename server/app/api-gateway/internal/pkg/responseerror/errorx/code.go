package errorx

type ErrorCode int32

//全局错误码
const (
	SERVER_COMMON_ERROR ErrorCode = 1001 + iota
	REUQEST_PARAM_ERROR
	GRPC_ERROR
	USERPASSWORDERROR
	UNAUTHORIZATION
	USER_PERMISSION_REJECT
	JWT_TOKEN_GENERATE_ERROR
	JWT_TOKEN_PARAM_GET_NERROR
	DB_ERROR
	DB_NOTFOUND
	DB_UPDATE_AFFECTED_ZERO_ERROR

	Validate_RegisterDefaultTranslations_ERROR
)

var message = map[ErrorCode]string{
	SERVER_COMMON_ERROR:                        "服务内部错误",
	REUQEST_PARAM_ERROR:                        "请求参数错误",
	GRPC_ERROR:                                 "GRPC错误",
	USERPASSWORDERROR:                          "用户名密码错误",
	UNAUTHORIZATION:                            "认证错误",
	USER_PERMISSION_REJECT:                     "用户权限被拒绝",
	JWT_TOKEN_PARAM_GET_NERROR:                 "JWT TOKEN参数获取失败",
	JWT_TOKEN_GENERATE_ERROR:                   "JWT TOKEN生成失败",
	DB_ERROR:                                   "数据库错误",
	DB_NOTFOUND:                                "数据库无记录",
	DB_UPDATE_AFFECTED_ZERO_ERROR:              "更新数据影响行数为0",
	Validate_RegisterDefaultTranslations_ERROR: "注册默认翻译失败",
}

func ErrorxMessage(errorxcode ErrorCode) string {
	if msg, ok := message[errorxcode]; ok {
		return msg
	} else {
		return "错误未定义"
	}
}
