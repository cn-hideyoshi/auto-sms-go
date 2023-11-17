package rpc

import (
	companyV1 "blog.hideyoshi.top/common/pkg/service/company.v1"
	"blog.hideyoshi.top/company/internal/handler"
	"context"
)

type CompanyLoginService struct {
	companyV1.UnimplementedCompanyLoginServiceServer
}

func NewCompanyLoginService() *CompanyLoginService {
	return &CompanyLoginService{}
}

func (*CompanyLoginService) Login(ctx context.Context, req *companyV1.CompanyLoginRequest) (*companyV1.CompanyLoginResponse, error) {
	loginHandler := handler.CompanyLoginHandler{}
	return loginHandler.Login(req)
}

func (*CompanyLoginService) Register(ctx context.Context, req *companyV1.CompanyRegisterRequest) (*companyV1.CompanyRegisterResponse, error) {
	loginHandler := handler.CompanyLoginHandler{}
	return loginHandler.Register(req)
}
