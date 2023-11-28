package main

import (
	"blog.hideyoshi.top/gateway/config"
	_ "blog.hideyoshi.top/gateway/internal/api"
	"blog.hideyoshi.top/gateway/router"
	_ "blog.hideyoshi.top/gateway/rpc"
	"blog.hideyoshi.top/gateway/server"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	//go func() {
	//	register := server.NewRegister()
	//	defer register.Stop()
	//	server.RegisterGateway()
	//}()
	server.Run(r, config.C.Http.Name, config.C.Http.Addr)
}
