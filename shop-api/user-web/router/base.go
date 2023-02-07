package router

import (
	"github.com/gin-gonic/gin"

	"shop-api/user-web/api"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.GET("captcha", api.GetCaptchar)
		BaseRouter.POST("send_sms", api.SendAliSms)
	}
}
