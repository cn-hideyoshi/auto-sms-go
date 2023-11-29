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
	"time"
)

type HandlerCompanyInfo struct {
}

func (*HandlerCompanyInfo) GetCompanyInfo(ctx *gin.Context) {
	info, _ := ctx.Get("LoginInfo")
	resp := types.Result{}
	var getCompanyInfoReq companyV1.GetCompanyInfoRequest
	getCompanyInfoReq.CompanyId = info.(*companyV1.CompanyInfo).CompanyId
	rpcResp, err := company.GetCompanyInfo(ctx, &getCompanyInfoReq)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.FailMsg(err.Error()))
		return
	}

	c := &model.Company{}
	err = copier.Copy(c, rpcResp.CompanyInfo)
	c.CreateTime = time.Unix(rpcResp.CompanyInfo.CreateTime, 0)
	c.UpdateTime = time.Unix(rpcResp.CompanyInfo.UpdateTime, 0)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.Fail(ecode.AuthError))
		return
	}
	ctx.JSON(http.StatusOK, resp.Success(types.GetCompanyInfoResponse{
		CompanyInfo: c,
	}))
}

func (*HandlerCompanyInfo) UpdateCompanyInfo(ctx *gin.Context) {
	resp := types.Result{}
	var loginReq companyV1.UpdateCompanyInfoRequest

	err := ctx.Bind(&loginReq)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.Fail(ecode.ParamsError))
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
