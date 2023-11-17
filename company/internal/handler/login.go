package handler

import (
	"blog.hideyoshi.top/common/pkg/ecode"
	companyV1 "blog.hideyoshi.top/common/pkg/service/company.v1"
)

type CompanyLoginHandler struct {
	CompanyHandler
}

func (ch *CompanyHandler) Register(req *companyV1.CompanyRegisterRequest) (*companyV1.CompanyRegisterResponse, error) {

	return &companyV1.CompanyRegisterResponse{
		Response: &companyV1.CompanyResponse{
			Code: ecode.SUCCESS,
			Msg:  ecode.GetMsg(ecode.SUCCESS),
		},
	}, nil
}

func (ch *CompanyHandler) Login(req *companyV1.CompanyLoginRequest) (*companyV1.CompanyLoginResponse, error) {
	return &companyV1.CompanyLoginResponse{
		Response: &companyV1.CompanyResponse{
			Code: ecode.SUCCESS,
			Msg:  ecode.GetMsg(ecode.SUCCESS),
		},
		Token: "test",
	}, nil
}
