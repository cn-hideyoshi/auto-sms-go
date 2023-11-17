package company

import (
	"blog.hideyoshi.top/common/model"
	companyV1 "blog.hideyoshi.top/common/pkg/service/company.v1"
	"blog.hideyoshi.top/gateway/internal/service/company"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerCompanyLogin struct {
}

func (*HandlerCompanyLogin) Login(ctx *gin.Context) {
	resp := &model.Result{}
	var loginReq companyV1.CompanyLoginRequest
	err := ctx.Bind(&loginReq)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.Fail(500, "获取参数失败"))
		return
	}
	ctx.JSON(http.StatusOK, resp.Success("Login"))
}

func (*HandlerCompanyLogin) Register(ctx *gin.Context) {
	resp := &model.Result{}
	req := companyV1.CompanyRegisterRequest{}

	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusOK, resp.Fail(400, "绑定参数错误"))
		return
	}

	_, err := company.Register(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.Fail(300, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, resp.SuccessNoData())
}
