package errmsg

const (
	SUCCESS               = 200
	ERROR                 = 500
	FRONTEND_PARAMS_ERROR = 406
)

var codeMsg = map[int]string{
	SUCCESS:               "OK",
	ERROR:                 "Fail",
	FRONTEND_PARAMS_ERROR: "Frontend Param Error",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
