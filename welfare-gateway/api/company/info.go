package company

import "github.com/gin-gonic/gin"

type HandlerCompanyInfo struct {
}

func (*HandlerCompanyInfo) GetCompanyInfo(ctx *gin.Context) {
	ctx.JSON(200, "login")
}

func (*HandlerCompanyInfo) UpdateCompanyInfo(ctx *gin.Context) {
	ctx.JSON(200, "Register")
}

func (*HandlerCompanyInfo) DeleteCompanyInfo(ctx *gin.Context) {
	ctx.JSON(200, "Register")
}
