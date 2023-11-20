package company

import (
	"blog.hideyoshi.top/common/pkg/db/model"
	"blog.hideyoshi.top/common/pkg/ecode"
	companyV1 "blog.hideyoshi.top/common/pkg/service/company.v1"
	"blog.hideyoshi.top/common/types"
	"blog.hideyoshi.top/gateway/internal/service/company"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
)

type HandlerCompanyInfo struct {
}

func (*HandlerCompanyInfo) GetCompanyInfo(ctx *gin.Context) {
	resp := types.Result{}
	var CompanyInfo companyV1.CompanyInfo

	err := ctx.Bind(&CompanyInfo)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.Fail(ecode.PARAMS_ERROR))
		return
	}
	rpcResp, err := company.GetCompanyInfo(ctx, &CompanyInfo)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.FailMsg(err.Error()))
		return
	}

	c := &model.Company{}
	err = copier.Copy(c, rpcResp.CompanyInfo)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.Fail(ecode.ERROR))
		return
	}
	ctx.JSON(http.StatusOK, resp.Success(types.GetCompanyInfoResponse{
		CompanyInfo: c,
	}))
}

func (*HandlerCompanyInfo) UpdateCompanyInfo(ctx *gin.Context) {
	resp := types.Result{}
	var loginReq companyV1.CompanyInfo

	err := ctx.Bind(&loginReq)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.Fail(ecode.PARAMS_ERROR))
		return
	}
	_, err = company.UpdateCompanyInfo(ctx, &loginReq)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.FailMsg(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, resp.SuccessNoData())
}

func (*HandlerCompanyInfo) DeleteCompanyInfo(ctx *gin.Context) {
	ctx.JSON(200, "Register")
}
