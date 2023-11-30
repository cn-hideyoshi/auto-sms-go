package server

import (
	"blog.hideyoshi.top/company/router"
)

func Start() {
	register := router.RegisterEtcd()
	register.Stop()
	router.RegisterGrpc()
}
