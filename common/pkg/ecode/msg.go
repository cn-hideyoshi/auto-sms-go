package ecode

// codeMsg error code msg
var codeMsg = map[int32]string{
	//system
	SUCCESS:     "ok",
	ERROR:       "fail",
	ParamsError: "param error",
	AuthError:   "auth error",
	PanicErr:    "panic error",

	//company
	CompanyNoExists:      "company no exists",
	CompanyPasswordError: "company password error",
	CompanyExists:        "company exists",
	//user
	UserNoExists:      "user no exists",
	UserPasswordError: "user password error",
	UserExists:        "user exists",
}

func GetMsg(key int32) string {
	msg, ok := codeMsg[key]
	if ok {
		return msg
	}
	return codeMsg[ERROR]
}
