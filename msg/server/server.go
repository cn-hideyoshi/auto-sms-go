package server

import "blog.hideyoshi.top/msg/router"

func Start() {
	register := router.RegisterEtcd()
	register.Stop()
	router.RegisterGrpc()
}
