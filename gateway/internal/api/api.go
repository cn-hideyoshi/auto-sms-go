package api

import (
	"blog.hideyoshi.top/gateway/internal/api/company"
	"blog.hideyoshi.top/gateway/internal/api/company_login"
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
	{
		h := &company_login.HandlerCompanyLogin{}
		api.POST("login", h.Login)
		api.POST("register", h.Register)
	}

	{
		h := &company.HandlerCompanyInfo{}
		api.GET("getCompanyInfo", h.GetCompanyInfo)
		api.PUT("UpdateCompanyInfo", h.UpdateCompanyInfo)
		api.POST("deleteCompanyInfo", h.DeleteCompanyInfo)
	}
}
