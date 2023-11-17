package server

import "blog.hideyoshi.top/user/router"

func Start() {
	register := router.RegisterEtcd()
	register.Stop()
	router.RegisterGrpc()
}
