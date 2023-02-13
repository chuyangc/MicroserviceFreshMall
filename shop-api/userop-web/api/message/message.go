package message

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
)

// List 获取留言信息
func List(c *gin.Context) {
	// 构建服务请求对象
	request := proto.MessageRequest{}

	// 获取当前登录用户
	claims, _ := c.Get("claims")
	model := claims.(*models.CustomClaims)

	// 判断当前用户是否是普通用户
	// TODO AuthorityId的值可能不正确
	if model.AuthorityId == 1 {
		request.UserId = int32(model.ID)
	}

	// 向服务发起请求
	rsp, err := global.MessageClient.MessageList(context.Background(), &request)
	if err != nil {
		zap.S().Errorw("[List] 查询 [留言信息列表]失败")
		api.HandleGrpcErrorToHttp(err, c)
		return
	}

	// 构建返回数据格式
	reMap := gin.H{
		"total": rsp.Total,
	}
	messageList := make([]interface{}, 0)
	for _, message := range rsp.Data {
		temp := gin.H{
			"id":          message.Id,
			"userId":      message.UserId,
			"messageType": message.MessageType,
			"subject":     message.Subject,
			"message":     message.Message,
			"file":        message.File,
		}
		messageList = append(messageList, temp)
	}
	reMap["data"] = messageList
	c.JSON(http.StatusOK, reMap)
}

// New 新增留言信息
func New(c *gin.Context) {
	// 校验form表单数据
	messageForm := forms.MessageForm{}
	if err := c.ShouldBindJSON(&messageForm); err != nil {
		api.HandleValidatorError(c, err)
		return
	}

	// 获取当前登录用户id
	userId, _ := c.Get("userId")

	// 构建服务请求对象
	request := proto.MessageRequest{
		UserId:      int32(userId.(uint)),
		MessageType: messageForm.MessageType,
		Subject:     messageForm.Subject,
		Message:     messageForm.Message,
		File:        messageForm.File,
	}

	// 向服务发起请求
	rsp, err := global.MessageClient.CreateMessage(context.Background(), &request)
	if err != nil {
		zap.S().Errorw("[New] 新建[留言信息]失败")
		api.HandleGrpcErrorToHttp(err, c)
		return
	}

	// 构建返回数据格式
	c.JSON(http.StatusOK, gin.H{
		"id": rsp.Id,
	})
}
