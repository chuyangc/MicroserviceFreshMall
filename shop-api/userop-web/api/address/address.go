package address

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

// List 获取收货地址信息
func List(c *gin.Context) {
	// 构建服务请求对象
	request := proto.AddressRequest{}

	// 获取当前登录用户信息
	claims, _ := c.Get("claims")
	currentUser := claims.(*models.CustomClaims)

	// 判断登录用户是否是普通用户
	// TODO 此处的AuthorityId的值可能不正确
	if currentUser.AuthorityId == 1 {
		request.UserId = int32(currentUser.ID)
	}

	// 向服务发起请求
	rsp, err := global.AddressClient.GetAddressList(context.Background(), &request)
	if err != nil {
		zap.S().Errorw("获取地址列表失败")
		api.HandleGrpcErrorToHttp(err, c)
		return
	}

	// 构建返回数据格式
	reMap := gin.H{
		"total": rsp.Total,
	}

	addressList := make([]interface{}, 0)
	for _, address := range rsp.Data {
		temp := gin.H{
			"id":            address.Id,
			"user_id":       address.UserId,
			"province":      address.Province,
			"city":          address.City,
			"district":      address.District,
			"address":       address.Address,
			"signer_name":   address.SignerName,
			"signer_mobile": address.SignerMobile,
		}
		addressList = append(addressList, temp)
	}

	reMap["data"] = addressList

	c.JSON(http.StatusOK, reMap)
}

// New 新建收货地址
func New(c *gin.Context) {
	// form表单校验
	createAddressForm := forms.AddressForm{}
	if err := c.ShouldBindJSON(&createAddressForm); err != nil {
		api.HandleValidatorError(c, err)
		return
	}

	// 获取当前登录用户id
	userId, _ := c.Get("userId")

	// 向服务发起请求
	rsp, err := global.AddressClient.CreateAddress(context.Background(), &proto.AddressRequest{
		UserId:       int32(userId.(uint)),
		Province:     createAddressForm.Province,
		City:         createAddressForm.City,
		District:     createAddressForm.District,
		Address:      createAddressForm.Address,
		SignerName:   createAddressForm.SignerName,
		SignerMobile: createAddressForm.SignerMobile,
	})

	if err != nil {
		zap.S().Errorw("新建地址失败")
		api.HandleGrpcErrorToHttp(err, c)
		return
	}

	// 构建返回数据格式
	c.JSON(http.StatusOK, gin.H{
		"id": rsp.Id,
	})
}

// Delete 删除收货地址
func Delete(c *gin.Context) {
	// 获取当前需要操作的数据id
	id := c.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	// 向服务发起请求
	_, err = global.AddressClient.DeleteAddress(context.Background(), &proto.AddressRequest{Id: int32(i)})
	if err != nil {
		zap.S().Errorw("删除地址失败")
		api.HandleGrpcErrorToHttp(err, c)
		return
	}

	// 构建返回数据格式
	c.JSON(http.StatusOK, gin.H{
		"msg": "删除成功",
	})
}

// Update 修改收货地址
func Update(c *gin.Context) {
	// form表单校验
	updateAddressForm := forms.AddressForm{}
	if err := c.ShouldBindJSON(&updateAddressForm); err != nil {
		api.HandleValidatorError(c, err)
		return
	}

	// 获取当前需要操作的数据id
	id := c.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	// 向服务发起请求
	_, err = global.AddressClient.UpdateAddress(context.Background(), &proto.AddressRequest{
		Id:           int32(i),
		Province:     updateAddressForm.Province,
		City:         updateAddressForm.City,
		District:     updateAddressForm.District,
		Address:      updateAddressForm.Address,
		SignerName:   updateAddressForm.SignerName,
		SignerMobile: updateAddressForm.SignerMobile,
	})
	if err != nil {
		zap.S().Errorw("更新地址失败")
		api.HandleGrpcErrorToHttp(err, c)
		return
	}
	// 返回数据
	c.Status(http.StatusOK)
}
