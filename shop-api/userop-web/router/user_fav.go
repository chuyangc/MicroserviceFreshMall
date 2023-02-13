package router

import (
	"github.com/gin-gonic/gin"
	"shop-api/userop-web/api/user_fav"
	"shop-api/userop-web/middlewares"
)

func InitUserFavRouter(Router *gin.RouterGroup) {
	UserFavRouter := Router.Group("userfavs")
	{
		UserFavRouter.DELETE("/:id", middlewares.JWTAuth(), user_fav.Delete)
		UserFavRouter.GET("/:id", middlewares.JWTAuth(), user_fav.Detail)
		UserFavRouter.POST("", middlewares.JWTAuth(), user_fav.New)
		UserFavRouter.GET("", middlewares.JWTAuth(), user_fav.List)
	}
}
