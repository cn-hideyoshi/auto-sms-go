package server

import "blog.hideyoshi.top/welfare/router"

func Start() {
	register := router.RegisterEtcd()
	register.Stop()
	router.RegisterGrpc()
}
