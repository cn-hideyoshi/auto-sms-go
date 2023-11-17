package api

import (
	"blog.hideyoshi.top/gateway/internal/api/company"
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
		h := &company.HandlerCompanyLogin{}
		api.POST("Login", h.Login)
		api.POST("Register", h.Register)
	}

	{
		h := &company.HandlerCompanyInfo{}
		api.GET("GetCompanyInfo", h.GetCompanyInfo)
		api.PUT("UpdateCompanyInfo", h.UpdateCompanyInfo)
		api.POST("DeleteCompanyInfo", h.DeleteCompanyInfo)
	}
}
