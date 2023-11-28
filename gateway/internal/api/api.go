package api

import (
	"blog.hideyoshi.top/gateway/internal/api/company"
	"blog.hideyoshi.top/gateway/internal/api/user"
	"blog.hideyoshi.top/gateway/internal/middleware"
	"blog.hideyoshi.top/gateway/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	router.Register(&RouterCompany{})
}

type RouterCompany struct {
}

func (*RouterCompany) Router(r *gin.Engine) {
	api := r.Group("api")
	companyApi := api.Group("company")
	h := &company.HandlerCompanyLogin{}
	companyApi.POST("T", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"A": 2})
	})
	companyApi.POST("Login", h.Login)
	//companyApi.POST("Register", h.Register)

	authCompany := api.Group("/company")
	{
		authCompany.Use(middleware.AuthMiddleware())
		h := &company.HandlerCompanyInfo{}
		authCompany.GET("GetCompanyInfo", h.GetCompanyInfo)
		authCompany.PUT("UpdateCompanyInfo", h.UpdateCompanyInfo)
		authCompany.POST("DeleteCompanyInfo", h.DeleteCompanyInfo)
	}

	userApi := api.Group("/user")
	{
		h := &user.HandlerUserLogin{}
		userApi.POST("Login", h.Login)
	}

}
