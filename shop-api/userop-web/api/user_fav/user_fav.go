package user_fav

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"shop-api/userop-web/api"
	"shop-api/userop-web/forms"
	"shop-api/userop-web/global"
	"shop-api/userop-web/models"
	"shop-api/userop-web/proto"
	"strconv"
)

// List 获取用户收藏信息列表
func List(c *gin.Context) {
	// 构建服务请求对象
	request := proto.UserFavRequest{}

	// 获取当前登录用户信息
	claims, _ := c.Get("claims")
	currentUser := claims.(*models.CustomClaims)

	// 判断当前用户是否是普通用户
	// TODO 此处AuthorityId的值可能不正确
	if currentUser.AuthorityId == 1 {
		request.UserId = int32(currentUser.ID)
	}

	// 向服务发起请求
	userFavRsp, err := global.UserFavClient.GetFavList(context.Background(), &request)
	if err != nil {
		zap.S().Errorw("获取收藏列表失败")
		api.HandleGrpcErrorToHttp(err, c)
		return
	}

	ids := make([]int32, 0)
	for _, item := range userFavRsp.Data {
		ids = append(ids, item.GoodsId)
	}

	if len(ids) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"total": 0,
		})
		return
	}

	// 请求商品服务，获取商品信息返回
	goods, err := global.GoodsSrvClient.BatchGetGoods(context.Background(), &proto.BatchGoodsIdInfo{
		Id: ids,
	})
	if err != nil {
		zap.S().Errorw("[List] 批量查询[商品列表]失败")
		api.HandleGrpcErrorToHttp(err, c)
		return
	}

	// 构建返回数据格式
	reMap := map[string]interface{}{
		"total": userFavRsp.Total,
	}

	goodsList := make([]interface{}, 0)
	for _, item := range userFavRsp.Data {
		data := gin.H{
			"id": item.GoodsId,
		}

		for _, good := range goods.Data {
			if item.GoodsId == good.Id {
				data["name"] = good.Name
				data["shop_price"] = good.ShopPrice
			}
		}

		goodsList = append(goodsList, data)
	}
	reMap["data"] = goodsList
	c.JSON(http.StatusOK, reMap)
}

// New 新建用户收藏信息
func New(c *gin.Context) {
	// form表单校验
	userFavForm := forms.UserFavForm{}
	if err := c.ShouldBindJSON(&userFavForm); err != nil {
		api.HandleValidatorError(c, err)
		return
	}

	// 向商品服务发起请求，判断用户输入的商品id是否存在
	_, err := global.GoodsSrvClient.GetGoodsDetail(context.Background(), &proto.GoodInfoRequest{
		Id: userFavForm.GoodsId,
	})
	if err != nil {
		zap.S().Errorw("商品[%s]不存在")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "商品不存在",
		})
	}

	// 向用户操作服务发起请求，完成新建用户收藏操作
	userId, _ := c.Get("userId")
	_, err = global.UserFavClient.AddUserFav(context.Background(), &proto.UserFavRequest{
		UserId:  int32(userId.(uint)),
		GoodsId: userFavForm.GoodsId,
	})

	if err != nil {
		zap.S().Errorw("添加收藏记录失败")
		api.HandleGrpcErrorToHttp(err, c)
		return
	}
	// 构建返回数据格式
	c.JSON(http.StatusOK, gin.H{
		"msg": "收藏成功",
	})
}

// Delete 删除用户收藏信息
func Delete(c *gin.Context) {
	// 获取需要操作的数据id
	id := c.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	// 获取当前登录的用户
	userId, _ := c.Get("userId")

	// 向服务发起请求【这里不用向商品服务发起请求，如果没有商品id不存在，底层会处理】
	_, err = global.UserFavClient.DeleteUserFav(context.Background(), &proto.UserFavRequest{
		UserId:  int32(userId.(uint)),
		GoodsId: int32(i),
	})
	if err != nil {
		zap.S().Errorw("删除收藏记录失败")
		api.HandleGrpcErrorToHttp(err, c)
		return
	}
	// 构建返回数据格式
	c.JSON(http.StatusOK, gin.H{
		"msg": "删除成功",
	})
}

// Detail 获取用户收藏详情
func Detail(c *gin.Context) {
	// 获取需要操作的数据id
	goodsId := c.Param("id")
	goodsIdInt, err := strconv.ParseInt(goodsId, 10, 32)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	userId, _ := c.Get("userId")
	_, err = global.UserFavClient.GetUserFavDetail(context.Background(), &proto.UserFavRequest{
		UserId:  int32(userId.(uint)),
		GoodsId: int32(goodsIdInt),
	})
	if err != nil {
		zap.S().Errorw("查询收藏状态失败")
		api.HandleGrpcErrorToHttp(err, c)
		return
	}

	c.Status(http.StatusOK)
}
