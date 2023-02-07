package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"shop-api/user-web/middlewares"
	"shop-api/user-web/router"
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
	ApiGroup := Router.Group("/u/v1")
	// 初始化用户路由
	router.InitUserRouter(ApiGroup)
	// 初始化基础路由
	router.InitBaseRouter(ApiGroup)

	return Router
}
