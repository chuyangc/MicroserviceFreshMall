package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"shop-api/userop-web/middlewares"
	"shop-api/userop-web/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	//配置服务发现和注册中心
	Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})

	// 配置跨域解决方案
	Router.Use(middlewares.Cors())
	ApiGroup := Router.Group("/up/v1")
	// 初始化商品路由
	router.InitAddressRouter(ApiGroup)
	router.InitMessageRouter(ApiGroup)
	router.InitUserFavRouter(ApiGroup)

	return Router
}
