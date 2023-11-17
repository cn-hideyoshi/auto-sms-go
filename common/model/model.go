package model

import "blog.hideyoshi.top/common/pkg/ecode"

type Result struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func (r *Result) Success(data any) *Result {
	r.Code = ecode.SUCCESS
	r.Message = ecode.GetMsg(r.Code)
	r.Data = data
	return r
}

func (r *Result) SuccessNoData() *Result {
	r.Success("")
	return r
}

func (r *Result) Fail(code int) *Result {
	r.Code = code
	r.Message = ecode.GetMsg(code)
	return r
}

func (r *Result) FailMsg(message string) *Result {
	r.Code = ecode.PANIC_ERR
	r.Message = message
	return r
}
