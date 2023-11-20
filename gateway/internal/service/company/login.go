package company

import (
	companyV1 "blog.hideyoshi.top/common/pkg/service/company.v1"
	"blog.hideyoshi.top/gateway/rpc"
	"context"
	"errors"
	"log"
)

func Login(ctx context.Context, req *companyV1.CompanyLoginRequest) (*companyV1.CompanyLoginResponse, error) {
	resp, err := rpc.Server.CompanyLoginClient.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Response.Code != 200 {
		log.Println("rpc code err", resp)
		err := errors.New(resp.Response.Msg)
		return nil, err
	}
	return resp, nil
}

func Register(ctx context.Context, req *companyV1.CompanyRegisterRequest) (*companyV1.CompanyRegisterResponse, error) {
	resp, err := rpc.Server.CompanyLoginClient.Register(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Response.Code != 200 {
		err := errors.New(resp.Response.Msg)
		return nil, err
	}
	return resp, nil
}
