package server

import (
	companyV1 "blog.hideyoshi.top/common/pkg/service/company.v1"
	"blog.hideyoshi.top/common/server"
	"blog.hideyoshi.top/company/config"
	"blog.hideyoshi.top/company/pkg/rpc"
	"google.golang.org/grpc"
)

func Start() {
	grpcServer := server.NewGrpcServer(config.C.Etcd, config.C.Grpc)
	grpcServer.RegisterEtcd()
	c := server.GrpcConfig{
		Addr: config.C.Grpc.Addr,
		RegisterFunc: func(server *grpc.Server) {
			companyV1.RegisterCompanyLoginServiceServer(server, rpc.NewCompanyLoginService())
			companyV1.RegisterCompanyInfoServiceServer(server, rpc.NewCompanyInfoService())
			companyV1.RegisterDepartmentInfoServiceServer(server, rpc.NewDepartmentInfoService())
		},
	}
	grpcServer.RunGrpc(c)
}
