package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"shop-api/goods-web/middlewares"
	"shop-api/goods-web/router"
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
	ApiGroup := Router.Group("/g/v1")
	// 初始化商品路由
	router.InitGoodsRouter(ApiGroup)
	router.InitCategoryRouter(ApiGroup)
	router.InitBannerRouter(ApiGroup)
	router.InitBrandRouter(ApiGroup)

	return Router
}
