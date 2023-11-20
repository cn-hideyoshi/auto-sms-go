package user

import (
	"blog.hideyoshi.top/common/pkg/ecode"
	userV1 "blog.hideyoshi.top/common/pkg/service/user.v1"
	"blog.hideyoshi.top/common/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerUserLogin struct {
}

func (*HandlerUserLogin) Login(ctx *gin.Context) {
	resp := &types.Result{}
	var loginReq userV1.UserLoginRequest
	err := ctx.Bind(&loginReq)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.Fail(ecode.PARAMS_ERROR))
		return
	}
	ctx.JSON(http.StatusOK, resp.Success("login"))
}

func (*HandlerUserLogin) Register(ctx *gin.Context) {
	resp := &types.Result{}
	ctx.JSON(http.StatusOK, resp.Success("Register"))
}
