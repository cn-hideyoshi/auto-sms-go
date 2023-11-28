package server

import "blog.hideyoshi.top/msg/router"

func Start() {
	register := router.RegisterEtcd()
	defer register.Stop()
	router.RegisterGrpc()
}
