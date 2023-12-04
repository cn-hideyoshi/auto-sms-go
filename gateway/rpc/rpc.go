package rpc

import (
	"blog.hideyoshi.top/common/pkg/rpc"
	"blog.hideyoshi.top/gateway/config"
	"log"
)

// Server is a singleton instance of GatewayServer.
var Server *rpc.CommonDiscoveryServer

// init initializes the GatewayServer and gRPC clients.
func init() {
	Server = rpc.NewCommonDiscoveryServer(config.C.Etcd)
	Server.NewRpcClient("user", &Server.UserLoginClient)
	Server.NewRpcClient("company", &Server.CompanyLoginClient)
	Server.NewRpcClient("company", &Server.CompanyInfoClient)
	Server.NewRpcClient("company", &Server.DepartmentInfoClient)
	//defer Server.ResolverClose()
	log.Println("init client success...")
}
