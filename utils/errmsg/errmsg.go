package errmsg

const (
	SUCCESS                = 200
	ERROR                  = 500
	FRONTEND_PARAMS_ERROR  = 406
	USER_EMAIL_EXIST       = 407
	PHONE_NUMBER_NOT_VALID = "电话号码必须是以84开头的11位数"
	PASSWORD_NOT_VALID     = "密码必须包含一个特殊字符"
)

var codeMsg = map[int]string{
	SUCCESS:               "OK",
	ERROR:                 "Fail",
	FRONTEND_PARAMS_ERROR: "页面参数错误",
	USER_EMAIL_EXIST:      "用户邮箱已经存在",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
