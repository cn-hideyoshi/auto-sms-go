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
	server.Run(r, config.C.Server.Name, config.C.Server.Addr)
}
