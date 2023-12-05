package service

import (
	loginServiceV1 "blog.hideyoshi.top/common/pkg/service/user.v1"
	"blog.hideyoshi.top/user/internal/handler"
	"context"
)

type UserPhoneService struct {
	loginServiceV1.UnimplementedUserPhoneServiceServer
}

func NewUserPhoneService() *UserPhoneService {
	return &UserPhoneService{}
}

func (*UserPhoneService) CreateUserPhone(ctx context.Context, req *loginServiceV1.CreateUserPhoneRequest) (*loginServiceV1.CreateUserPhoneResponse, error) {
	h := handler.PhoneHandler{}
	return h.CreateUserPhone(req), nil
}
func (*UserPhoneService) GetUserPhone(ctx context.Context, req *loginServiceV1.GetUserPhoneRequest) (*loginServiceV1.GetUserPhoneResponse, error) {
	h := handler.PhoneHandler{}
	return h.GetUserPhone(req), nil
}
