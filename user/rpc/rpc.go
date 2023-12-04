package rpc

import (
	"blog.hideyoshi.top/common/pkg/rpc"
	"blog.hideyoshi.top/user/config"
	"log"
)

// Clients is a singleton instance of GatewayServer.
var Clients *rpc.CommonDiscoveryServer

// init initializes the UserServer and gRPC clients.
func init() {
	Clients = rpc.NewCommonDiscoveryServer(config.C.Etcd)
	Clients.NewRpcClient("company", &Clients.CompanyLoginClient)
	Clients.NewRpcClient("company", &Clients.CompanyInfoClient)
	Clients.NewRpcClient("company", &Clients.DepartmentInfoClient)
	//defer Clients.ResolverClose()
	log.Println("init client success...")
}
