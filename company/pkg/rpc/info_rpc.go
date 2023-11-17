package rpc

import (
	loginServiceV1 "blog.hideyoshi.top/common/pkg/service/company.v1"
	"context"
)

type CompanyInfoService struct {
	loginServiceV1.UnimplementedCompanyInfoServiceServer
}

func NewCompanyInfoService() *CompanyInfoService {
	return &CompanyInfoService{}
}

func (*CompanyInfoService) GetCompanyInfo(ctx context.Context, req *loginServiceV1.CompanyInfoRequest) (*loginServiceV1.CompanyInfoResponse, error) {
	l := &loginServiceV1.CompanyInfoResponse{
		Test: "Admin",
	}

	return l, nil
}
