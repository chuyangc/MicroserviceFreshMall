package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"net/http"
)

// 声明变量用于验证码保存
var store = base64Captcha.DefaultMemStore

// GetCaptchar -> 获取验证码
func GetCaptchar(ctx *gin.Context) {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, pic, err := cp.Generate()
	if err != nil {
		zap.S().Errorf("生成验证码错误: %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成验证码错误",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"captchaId": id,
		"picPath":   pic,
	})
}
