package router

import (
	"github.com/gin-gonic/gin"
	"shop-api/goods-web/api/goods"
	"shop-api/goods-web/middlewares"
)

func InitGoodsRouter(Router *gin.RouterGroup) {
	GoodsRouter := Router.Group("goods")
	// zap.S().Info("配置商品相关的Url")
	{
		GoodsRouter.GET("", goods.List) //商品列表
		GoodsRouter.POST("", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.Create)
		GoodsRouter.GET("/:id", goods.Detail)                                                           //获取商品的详情
		GoodsRouter.DELETE("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.Delete)      //删除商品
		GoodsRouter.GET("/:id/stocks", goods.Stocks)                                                    //获取商品的库存
		GoodsRouter.PUT("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.Update)         //更新商品信息
		GoodsRouter.PATCH("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.UpdateStatus) //更新商品状态
	}
}
