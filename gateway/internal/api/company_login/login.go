package company_login

import (
	"blog.hideyoshi.top/common/model"
	login_service_v1 "blog.hideyoshi.top/gateway/pkg/service/login.service.v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerCompanyLogin struct {
}

func (*HandlerCompanyLogin) Login(ctx *gin.Context) {
	resp := &model.Result{}
	var loginReq login_service_v1.LoginMessage
	err := ctx.Bind(&loginReq)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.Fail(200, "获取参数失败"))
		return
	}

	//userService := ctx.Keys["us"]

	ctx.JSON(http.StatusOK, resp.Success("login"))
}

func (*HandlerCompanyLogin) Register(ctx *gin.Context) {
	resp := &model.Result{}
	ctx.JSON(http.StatusOK, resp.Success("Register"))
}
