package util

import (
	"blog.hideyoshi.top/common/pkg/ecode"
	companyV1 "blog.hideyoshi.top/common/pkg/service/company.v1"
)

func SetErrors(resp *companyV1.CompanyResponse, code int32) {
	resp.Code = code
	resp.Msg = ecode.GetMsg(code)
}
