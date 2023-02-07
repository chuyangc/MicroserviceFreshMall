package router

import (
	"github.com/gin-gonic/gin"
	"shop-api/user-web/api"
	"shop-api/user-web/middlewares"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	// zap.S().Info("配置用户相关的Url")
	{
		UserRouter.GET("list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
		UserRouter.POST("pwd_login", api.PassWordLogin)
		UserRouter.POST("register", api.Register)
	}
}
