package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
	"strings"
	"time"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"

	"shop-api/user-web/forms"
	"shop-api/user-web/global"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func GenerateSmsCode(width int) string {
	// 生成长度为width的随机数
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	numLen := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		_, _ = fmt.Fprintf(&sb, "%d", numeric[rand.Intn(numLen)])
	}
	return sb.String()
}

func SendAliSms(ctx *gin.Context) {
	//表单验证
	sendSmsForm := forms.SendSmsForm{}
	if _err := ctx.ShouldBind(&sendSmsForm); _err != nil {
		HandleValidatorError(ctx, _err)
		return
	}

	// 工程代码泄露可能会导致AccessKey泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html
	client, _err := CreateClient(tea.String(global.ServerConfig.AliSmsInfo.ApiKey),
		tea.String(global.ServerConfig.AliSmsInfo.ApiSecrect))
	if _err != nil {
		zap.S().Warnf("错误: %v", _err)
		//panic(_err)
	}

	smsCode := GenerateSmsCode(6)
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String("欧楚杨的博客"),
		TemplateCode:  tea.String("SMS_269045571"),
		PhoneNumbers:  tea.String(sendSmsForm.Mobile),
		TemplateParam: tea.String("{\"code\":\"" + smsCode + "\"}"),
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		_, _err = client.SendSmsWithOptions(sendSmsRequest, runtime)
		//将验证码保存 -> redis
		rdb := redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port),
		})
		rdb.Set(context.Background(), sendSmsForm.Mobile, smsCode, time.Duration(global.ServerConfig.RedisInfo.Expire)*time.Second)
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "发送成功",
		})
		if _err != nil {
			zap.S().Warnf("错误: %v", _err)
			//panic(_err)
		}
		return
	}()

	if tryErr != nil {
		var _error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			_error = _t
		} else {
			_error.Message = tea.String(tryErr.Error())
		}
		// 如有需要，请打印 error
		_, _err = util.AssertAsString(_error.Message)
		if _err != nil {
			zap.S().Warnf("错误: %v", _err)
		}
	}
	return
}
