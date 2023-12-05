package ecode

// system const
const (
	SUCCESS     = 200
	ERROR       = 500
	ParamsError = 400
	AuthError   = 401
	PanicErr    = 999
	JwtErr      = 595 + iota
	RedisErr
)

// company const
const (
	CompanyNoExists = 10001 + iota
	CompanyPasswordError
	CompanyExists
)

// user const
const (
	UserNoExists = 20001 + iota
	UserPasswordError
	UserExists
)

// msg const
const (
	MsgGroupNoExists = 30001 + iota
	MsgGroupExists
)
