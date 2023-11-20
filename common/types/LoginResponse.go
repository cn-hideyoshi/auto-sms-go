package types

import "blog.hideyoshi.top/common/pkg/db/model"

type CompanyLoginResponse struct {
	Token string `json:"token"`
}

type GetCompanyInfoResponse struct {
	CompanyInfo *model.Company
}
