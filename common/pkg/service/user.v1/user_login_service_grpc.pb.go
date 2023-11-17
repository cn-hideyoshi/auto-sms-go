// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: user/user_login_service.proto

package user_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	UserLoginService_Login_FullMethodName = "/user.v1.UserLoginService/Login"
)

// UserLoginServiceClient is the client API for UserLoginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserLoginServiceClient interface {
	Login(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error)
}

type userLoginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserLoginServiceClient(cc grpc.ClientConnInterface) UserLoginServiceClient {
	return &userLoginServiceClient{cc}
}

func (c *userLoginServiceClient) Login(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error) {
	out := new(UserLoginResponse)
	err := c.cc.Invoke(ctx, UserLoginService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserLoginServiceServer is the server API for UserLoginService service.
// All implementations must embed UnimplementedUserLoginServiceServer
// for forward compatibility
type UserLoginServiceServer interface {
	Login(context.Context, *UserLoginRequest) (*UserLoginResponse, error)
	mustEmbedUnimplementedUserLoginServiceServer()
}

// UnimplementedUserLoginServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserLoginServiceServer struct {
}

func (UnimplementedUserLoginServiceServer) Login(context.Context, *UserLoginRequest) (*UserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserLoginServiceServer) mustEmbedUnimplementedUserLoginServiceServer() {}

// UnsafeUserLoginServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserLoginServiceServer will
// result in compilation errors.
type UnsafeUserLoginServiceServer interface {
	mustEmbedUnimplementedUserLoginServiceServer()
}

func RegisterUserLoginServiceServer(s grpc.ServiceRegistrar, srv UserLoginServiceServer) {
	s.RegisterService(&UserLoginService_ServiceDesc, srv)
}

func _UserLoginService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLoginServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserLoginService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLoginServiceServer).Login(ctx, req.(*UserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserLoginService_ServiceDesc is the grpc.ServiceDesc for UserLoginService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserLoginService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.v1.UserLoginService",
	HandlerType: (*UserLoginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _UserLoginService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user_login_service.proto",
}