package company

import (
	"blog.hideyoshi.top/gateway/internal/middleware"
	"github.com/gin-gonic/gin"
)

type RouterCompany struct {
}

func (*RouterCompany) Router(r *gin.Engine) {
	api := r.Group("api")
	loginHandler := &HandlerCompanyLogin{}
	infoHandler := &HandlerCompanyInfo{}

	//no auth
	noAuthCompany := api.Group("company")
	{
		noAuthCompany.POST("Login", loginHandler.Login)
		noAuthCompany.POST("Register", loginHandler.Register)
	}

	//authed api
	authedCompany := api.Group("company")
	authedCompany.Use(middleware.AuthMiddleware())
	{
		authedCompany.GET("GetCompanyInfo", infoHandler.GetCompanyInfo)
		authedCompany.PUT("UpdateCompanyInfo", infoHandler.UpdateCompanyInfo)
		authedCompany.POST("DeleteCompanyInfo", infoHandler.DeleteCompanyInfo)
	}

}
