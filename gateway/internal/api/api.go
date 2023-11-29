package api

import (
	"blog.hideyoshi.top/gateway/internal/api/company"
	"blog.hideyoshi.top/gateway/internal/api/user"

	"blog.hideyoshi.top/gateway/router"
)

func init() {
	router.Register(&company.RouterCompany{})
	router.Register(&user.RouterUser{})
}
