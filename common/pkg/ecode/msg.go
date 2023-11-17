package ecode

// codeMsg error code msg
var codeMsg = map[int]string{
	//system
	SUCCESS:      "ok",
	ERROR:        "fail",
	PARAMS_ERROR: "param error",
	PANIC_ERR:    "panic error",

	//company
	COMPANY_NO_EXISTS:      "company no exists",
	COMPANY_PASSWORD_ERROR: "company password error",
	COMPANY_EXISTS:         "company exists",
	//user
	USER_NO_EXISTS:      "user no exists",
	USER_PASSWORD_ERROR: "user password error",
	USER_EXISTS:         "user exists",
}

func GetMsg(key int) string {
	msg, ok := codeMsg[key]
	if ok {
		return msg
	}
	return codeMsg[ERROR]
}
