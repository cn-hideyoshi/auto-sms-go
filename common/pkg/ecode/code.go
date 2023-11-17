package ecode

// system const
const (
	SUCCESS      = 200
	ERROR        = 500
	PARAMS_ERROR = 400
	PANIC_ERR    = 999
)

// company const
const (
	COMPANY_NO_EXISTS = 10001 + iota
	COMPANY_PASSWORD_ERROR
	COMPANY_EXISTS
)

// user const
const (
	USER_NO_EXISTS = 20001 + iota
	USER_PASSWORD_ERROR
	USER_EXISTS
)
