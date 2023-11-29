package company

import (
	companyV1 "blog.hideyoshi.top/common/pkg/service/company.v1"
	"blog.hideyoshi.top/gateway/rpc"
	"context"
	"errors"
)

func GetCompanyInfo(ctx context.Context, req *companyV1.GetCompanyInfoRequest) (*companyV1.CompanyInfoResponse, error) {
	resp, err := rpc.Server.CompanyInfoClient.GetCompanyInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Response.Code != 200 {
		err := errors.New(resp.Response.Msg)
		return nil, err
	}
	return resp, nil
}

func UpdateCompanyInfo(ctx context.Context, req *companyV1.UpdateCompanyInfoRequest) (*companyV1.CompanyInfoResponse, error) {
	resp, err := rpc.Server.CompanyInfoClient.UpdateCompanyInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Response.Code != 200 {
		err := errors.New(resp.Response.Msg)
		return nil, err
	}
	return resp, nil
}
