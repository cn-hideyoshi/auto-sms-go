package types

import model "blog.hideyoshi.top/common/pkg/db/model/company"

type CompanyLoginResponse struct {
	Token string `json:"token"`
}

type GetCompanyInfoResponse struct {
	CompanyInfo *model.Company
}
