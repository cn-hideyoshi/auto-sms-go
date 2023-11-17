package company

import (
	"blog.hideyoshi.top/common/model"
	"blog.hideyoshi.top/common/pkg/ecode"
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
		ctx.JSON(http.StatusOK, resp.Fail(ecode.PARAMS_ERROR))
		return
	}
	rpcResp, err := company.Login(ctx, &loginReq)
	ctx.JSON(http.StatusOK, resp.Success(rpcResp.Token))
}

func (*HandlerCompanyLogin) Register(ctx *gin.Context) {
	resp := &model.Result{}
	req := companyV1.CompanyRegisterRequest{}

	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusOK, resp.Fail(ecode.PARAMS_ERROR))
		return
	}

	_, err := company.Register(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.FailMsg(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, resp.SuccessNoData())
}
