package company_login

import (
	"blog.hideyoshi.top/welfare-common/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerCompanyLogin struct {
}

func (*HandlerCompanyLogin) Login(ctx *gin.Context) {
	resp := &model.Result{}
	ctx.JSON(http.StatusOK, resp.Success("login"))
}

func (*HandlerCompanyLogin) Register(ctx *gin.Context) {
	resp := &model.Result{}
	ctx.JSON(http.StatusOK, resp.Success("Register"))
}
