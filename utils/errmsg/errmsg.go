package errmsg

const (
	SUCCESS                = 200
	ERROR                  = 500
	FRONTEND_PARAMS_ERROR  = 406
	PHONE_NUMBER_NOT_VALID = "电话号码必须是以84开头的11位数"
)

var codeMsg = map[int]string{
	SUCCESS:               "OK",
	ERROR:                 "Fail",
	FRONTEND_PARAMS_ERROR: "Frontend Param Error",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
