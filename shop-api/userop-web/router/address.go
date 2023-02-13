package router

import (
	"github.com/gin-gonic/gin"
	"shop-api/userop-web/api/address"
	"shop-api/userop-web/middlewares"
)

func InitAddressRouter(Router *gin.RouterGroup) {
	AddressRouter := Router.Group("address")
	{
		AddressRouter.GET("", middlewares.JWTAuth(), address.List)
		AddressRouter.DELETE("/:id", middlewares.JWTAuth(), address.Delete)
		AddressRouter.POST("", middlewares.JWTAuth(), address.New)
		AddressRouter.PATCH("/:id", middlewares.JWTAuth(), address.Update)
	}
}
