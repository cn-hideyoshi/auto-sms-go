package service

import (
	loginServiceV1 "blog.hideyoshi.top/common/pkg/service/user.v1"
	"blog.hideyoshi.top/user/internal/handler"
	"context"
)

type UserInfoService struct {
	loginServiceV1.UnimplementedUserInfoServiceServer
}

func NewUserInfoService() *UserInfoService {
	return &UserInfoService{}
}

func (*UserInfoService) CreateUserInfo(ctx context.Context, req *loginServiceV1.CreateUserInfoRequest) (*loginServiceV1.CreateUserInfoResponse, error) {
	h := handler.InfoHandler{}
	return h.CreateUserInfo(req), nil
}

func (*UserInfoService) UpdateUserInfo(ctx context.Context, req *loginServiceV1.UpdateUserInfoRequest) (*loginServiceV1.UpdateUserInfoResponse, error) {
	h := handler.InfoHandler{}
	return h.UpdateUserInfo(req), nil
}

func (*UserInfoService) GetUserInfo(ctx context.Context, req *loginServiceV1.GetUserInfoRequest) (*loginServiceV1.GetUserInfoResponse, error) {
	h := handler.InfoHandler{}
	return h.GetUserInfo(req), nil
}
