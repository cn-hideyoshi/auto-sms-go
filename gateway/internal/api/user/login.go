package user

import (
	userV1 "blog.hideyoshi.top/common/pkg/service/user.v1"
	"blog.hideyoshi.top/common/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerCompanyLogin struct {
}

func (*HandlerCompanyLogin) Login(ctx *gin.Context) {
	resp := &types.Result{}
	var loginReq userV1.UserLoginRequest
	err := ctx.Bind(&loginReq)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.Fail(200, "获取参数失败"))
		return
	}
	ctx.JSON(http.StatusOK, resp.Success("login"))
}

func (*HandlerCompanyLogin) Register(ctx *gin.Context) {
	resp := &types.Result{}
	ctx.JSON(http.StatusOK, resp.Success("Register"))
}
