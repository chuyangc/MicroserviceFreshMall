package order

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"shop-api/order-web/api"
	"shop-api/order-web/forms"
	"shop-api/order-web/global"
	"shop-api/order-web/models"
	"shop-api/order-web/proto"
	"strconv"
)

func List(c *gin.Context) {
	// 获取订单列表
	userId, _ := c.Get("userId")
	claims, _ := c.Get("claims")

	request := proto.OrderFilterRequest{}

	// 如果是管理员用户则返回所有的订单
	model := claims.(*models.CustomClaims)
	// TODO 此处AuthorityId 可能不正确
	zap.S().Infof("model.AuthorityId -> %v", model.AuthorityId)
	if model.AuthorityId == 1 {
		request.UserId = int32(userId.(uint))
	}

	// 分页
	pages := c.DefaultQuery("p", "0")
	pagesInt, _ := strconv.Atoi(pages)
	request.Pages = int32(pagesInt)

	perNums := c.DefaultQuery("pnum", "0")
	perNumsInt, _ := strconv.Atoi(perNums)
	request.PagePerNums = int32(perNumsInt)

	request.Pages = int32(pagesInt)
	request.PagePerNums = int32(perNumsInt)

	rsp, err := global.OrderSrvClient.OrderList(context.Background(), &request)
	if err != nil {
		zap.S().Errorw("获取订单列表失败")
		api.HandleGrpcErrorToHttp(err, c)
		return
	}

	/*
		{
			"total": ,
			"data":[
				{

				}
			]
		}
	*/
	reMap := gin.H{
		"total": rsp.Total,
	}
	orderList := make([]interface{}, 0)

	for _, item := range rsp.Data {
		tmpMap := map[string]interface{}{}

		tmpMap["id"] = item.Id
		tmpMap["status"] = item.Status
		tmpMap["pay_type"] = item.PayType
		tmpMap["user"] = item.UserId
		tmpMap["post"] = item.Post
		tmpMap["total"] = item.Total
		tmpMap["address"] = item.Address
		tmpMap["name"] = item.Name
		tmpMap["mobile"] = item.Mobile
		tmpMap["order_sn"] = item.OrderSn
		tmpMap["id"] = item.Id
		tmpMap["add_time"] = item.AddTime // 直接显示时间

		orderList = append(orderList, tmpMap)
	}
	reMap["data"] = orderList
	c.JSON(http.StatusOK, reMap)
}

func New(c *gin.Context) {
	orderForm := forms.CreateOrderForm{}
	if err := c.ShouldBindJSON(&orderForm); err != nil {
		api.HandleValidatorError(c, err)
	}
	userId, _ := c.Get("userId")
	rsp, err := global.OrderSrvClient.CreateOrder(context.WithValue(context.Background(), "ginContext", c), &proto.OrderRequest{
		UserId:  int32(userId.(uint)),
		Name:    orderForm.Name,
		Mobile:  orderForm.Mobile,
		Address: orderForm.Address,
		Post:    orderForm.Post,
	})
	if err != nil {
		zap.S().Errorw("新建订单失败")
		api.HandleGrpcErrorToHttp(err, c)
		return
	}
	// TODO 返回支付宝的支付url
	c.JSON(http.StatusOK, gin.H{
		"id": rsp.Id,
	})
}

func Detail(c *gin.Context) {
	id := c.Param("id")
	userId, _ := c.Get("userId")
	i, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "url格式出错",
		})
		return
	}

	//如果是管理员用户则返回所有的订单
	request := proto.OrderRequest{
		// 订单Id
		Id: int32(i),
	}
	claims, _ := c.Get("claims")
	model := claims.(*models.CustomClaims)
	// TODO 此处AuthorityId 可能不正确
	zap.S().Infof("model.AuthorityId -> %v", model.AuthorityId)
	if model.AuthorityId == 1 {
		request.UserId = int32(userId.(uint))
	}

	rsp, err := global.OrderSrvClient.OrderDetail(context.Background(), &request)
	if err != nil {
		zap.S().Errorw("获取订单详情失败")
		api.HandleGrpcErrorToHttp(err, c)
		return
	}
	reMap := gin.H{}
	reMap["id"] = rsp.OrderInfo.Id
	reMap["status"] = rsp.OrderInfo.Status
	reMap["user"] = rsp.OrderInfo.UserId
	reMap["post"] = rsp.OrderInfo.Post
	reMap["total"] = rsp.OrderInfo.Total
	reMap["address"] = rsp.OrderInfo.Address
	reMap["name"] = rsp.OrderInfo.Name
	reMap["mobile"] = rsp.OrderInfo.Mobile
	reMap["pay_type"] = rsp.OrderInfo.PayType
	reMap["order_sn"] = rsp.OrderInfo.OrderSn

	// 订单里的商品
	goodsList := make([]interface{}, 0)
	for _, item := range rsp.Data {
		tmpMap := gin.H{
			"id":    item.GoodsId,
			"name":  item.GoodsName,
			"image": item.GoodsImage,
			"price": item.GoodsPrice,
			"nums":  item.Nums,
		}

		goodsList = append(goodsList, tmpMap)
	}
	reMap["goods"] = goodsList

	c.JSON(http.StatusOK, reMap)
}
