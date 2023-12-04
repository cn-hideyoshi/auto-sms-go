package service

import (
	loginServiceV1 "blog.hideyoshi.top/common/pkg/service/user.v1"
	"context"
)

type UserLoginService struct {
	loginServiceV1.UnimplementedUserLoginServiceServer
}

func NewUserLoginService() *UserLoginService {
	return &UserLoginService{}
}

func (*UserLoginService) Login(ctx context.Context, req *loginServiceV1.UserLoginRequest) (*loginServiceV1.UserLoginResponse, error) {
	l := &loginServiceV1.UserLoginResponse{
		//Test: "Admin",
	}
	return l, nil
}
