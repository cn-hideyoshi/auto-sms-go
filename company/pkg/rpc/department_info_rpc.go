package rpc

import (
	loginServiceV1 "blog.hideyoshi.top/common/pkg/service/company.v1"
	"blog.hideyoshi.top/company/internal/handler"
	"context"
)

type DepartmentInfoService struct {
	loginServiceV1.UnimplementedDepartmentInfoServiceServer
}

func NewDepartmentInfoService() *DepartmentInfoService {
	return &DepartmentInfoService{}
}

func (*DepartmentInfoService) CreateDepartment(ctx context.Context, req *loginServiceV1.CreateDepartmentRequest) (*loginServiceV1.CreateDepartmentResponse, error) {
	departmentInfoHandler := handler.DepartmentInfoHandler{}
	return departmentInfoHandler.CreateDepartment(req), nil
}

func (*DepartmentInfoService) GetDepartment(ctx context.Context, req *loginServiceV1.GetDepartmentRequest) (*loginServiceV1.GetDepartmentResponse, error) {
	departmentInfoHandler := handler.DepartmentInfoHandler{}
	return departmentInfoHandler.GetCompanyInfo(req), nil
}

func (*DepartmentInfoService) UpdateDepartment(ctx context.Context, req *loginServiceV1.UpdateDepartmentRequest) (*loginServiceV1.UpdateDepartmentResponse, error) {
	departmentInfoHandler := handler.DepartmentInfoHandler{}
	return departmentInfoHandler.UpdateDepartment(req), nil
}

func (*DepartmentInfoService) GetDepartmentTree(ctx context.Context, req *loginServiceV1.GetDepartmentTreeRequest) (*loginServiceV1.GetDepartmentTreeResponse, error) {
	departmentInfoHandler := handler.DepartmentInfoHandler{}
	return departmentInfoHandler.GetDepartmentTree(req), nil
}
