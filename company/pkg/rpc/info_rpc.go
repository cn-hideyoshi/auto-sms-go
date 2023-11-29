package rpc

import (
	loginServiceV1 "blog.hideyoshi.top/common/pkg/service/company.v1"
	"blog.hideyoshi.top/company/internal/handler"
	"context"
)

type CompanyInfoService struct {
	loginServiceV1.UnimplementedCompanyInfoServiceServer
}

func NewCompanyInfoService() *CompanyInfoService {
	return &CompanyInfoService{}
}

func (*CompanyInfoService) GetCompanyInfo(ctx context.Context, req *loginServiceV1.GetCompanyInfoRequest) (*loginServiceV1.CompanyInfoResponse, error) {
	companyInfoHandler := handler.CompanyInfoHandler{}
	return companyInfoHandler.GetCompanyInfo(req)
}

func (*CompanyInfoService) UpdateCompanyInfo(ctx context.Context, req *loginServiceV1.UpdateCompanyInfoRequest) (*loginServiceV1.CompanyInfoResponse, error) {
	companyInfoHandler := handler.CompanyInfoHandler{}
	return companyInfoHandler.UpdateCompanyInfo(req)
}
