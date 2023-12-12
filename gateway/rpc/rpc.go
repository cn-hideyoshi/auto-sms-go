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
	Server.NewRpcClient("company", &Server.CompanyLoginClient)
	Server.NewRpcClient("company", &Server.CompanyInfoClient)
	Server.NewRpcClient("company", &Server.DepartmentInfoClient)
	Server.NewRpcClient("user", &Server.UserPhoneClient)
	Server.NewRpcClient("user", &Server.UserInfoClient)
	Server.NewRpcClient("msg", &Server.MsgGroupClient)
	//defer Server.ResolverClose()
	log.Println("init client success...")
}
