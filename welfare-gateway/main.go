package main

import (
	"blog.hideyoshi.top/welfare-common/server"
	_ "blog.hideyoshi.top/welfare-gateway/api"
	"blog.hideyoshi.top/welfare-gateway/config"
	"blog.hideyoshi.top/welfare-gateway/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	server.Run(r, config.C.SC.Name, config.C.SC.Addr)
}
