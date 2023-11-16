package router

import (
	"github.com/gin-gonic/gin"
)

type Router interface {
	Router(r *gin.Engine)
}

var routers []Router

func InitRouter(r *gin.Engine) {
	for _, router := range routers {
		router.Router(r)
	}

	//router := New()
	//router.Route(&login.RouterCompany{}, r)
}

func Register(ro ...Router) {
	routers = append(routers, ro...)
}
