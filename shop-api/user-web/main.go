package main

import (
	"fmt"
	"github.com/spf13/viper"
	"shop-api/user-web/utils"

	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	myvalidator "shop-api/user-web/validator"

	"shop-api/user-web/global"
	"shop-api/user-web/initialize"
)

func main() {
	//初始化logger
	initialize.InitLogger()

	//初始化配置文件
	initialize.InitNacosConfig()

	//初始化Router
	Router := initialize.Routers()

	//初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}

	//初始化srv的连接
	initialize.InitConsulLBSrvConn()

	// 动态端口获取
	viper.AutomaticEnv()
	debug := viper.GetBool("SHOP_DEBUG")
	// 如果是本地开发环境端口号固定，线上环境启动获取动态空闲端口号
	if debug {
		port, err := utils.GetFreePort()
		if err == nil {
			global.ServerConfig.Port = port
		}
	}

	//注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", myvalidator.ValidateMobile)
		_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 非法的手机号码!", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}

	/*
		S()可以获取一个全局的sugar，可以设置全局logger
		S和L函数提供一个安全方式访问
	*/
	zap.S().Debugf("启动服务器, 端口: %d", global.ServerConfig.Port)
	if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("启动失败: ", err.Error())
	}
}
