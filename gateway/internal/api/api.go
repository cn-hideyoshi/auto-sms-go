package api

import (
	"blog.hideyoshi.top/gateway/internal/api/company"
	"blog.hideyoshi.top/gateway/internal/api/user"
	"blog.hideyoshi.top/gateway/router"
	"github.com/gin-gonic/gin"
)

func init() {
	router.Register(&RouterCompany{})
}

type RouterCompany struct {
}

func (*RouterCompany) Router(r *gin.Engine) {
	api := r.Group("/api")
	companyApi := api.Group("/company")
	{
		h := &company.HandlerCompanyLogin{}
		companyApi.POST("Login", h.Login)
		companyApi.POST("Register", h.Register)
	}
	{
		h := &company.HandlerCompanyInfo{}
		companyApi.GET("GetCompanyInfo", h.GetCompanyInfo)
		companyApi.PUT("UpdateCompanyInfo", h.UpdateCompanyInfo)
		companyApi.POST("DeleteCompanyInfo", h.DeleteCompanyInfo)
	}

	userApi := api.Group("/user")
	{
		h := &user.HandlerUserLogin{}
		userApi.POST("Login", h.Login)
		userApi.POST("Register", h.Register)
	}
}
