package model

type Result struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func (r *Result) Success(data any) *Result {
	r.Code = 200
	r.Message = "success"
	r.Data = data
	return r
}

func (r *Result) SuccessNoData() *Result {
	r.Success("")
	return r
}

func (r *Result) Fail(code int, message string) *Result {
	r.Code = code
	r.Message = message
	return r
}
