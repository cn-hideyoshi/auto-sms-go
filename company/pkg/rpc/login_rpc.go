package rpc

import (
	companyV1 "blog.hideyoshi.top/common/pkg/service/company.v1"
	"context"
)

type CompanyLoginService struct {
	companyV1.UnimplementedCompanyLoginServiceServer
}

func NewCompanyLoginService() *CompanyLoginService {
	return &CompanyLoginService{}
}

func (*CompanyLoginService) Login(ctx context.Context, req *companyV1.CompanyLoginRequest) (*companyV1.CompanyLoginResponse, error) {
	l := &companyV1.CompanyLoginResponse{
		Response: &companyV1.CompanyResponse{
			Code: 200,
			Msg:  "Admin",
		},
	}
	return l, nil
}

func (*CompanyLoginService) Register(ctx context.Context, req *companyV1.CompanyRegisterRequest) (*companyV1.CompanyRegisterResponse, error) {
	l := &companyV1.CompanyRegisterResponse{
		Response: &companyV1.CompanyResponse{
			Code: 200,
			Msg:  "Admin",
		},
	}
	return l, nil
}
