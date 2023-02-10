package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"shop-api/oss-web/utils"
	"shop-api/oss-web/utils/register/consul"
	"syscall"

	"go.uber.org/zap"
	"shop-api/oss-web/global"
	"shop-api/oss-web/initialize"
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

	/*
		S()可以获取一个全局的sugar，可以设置全局logger
		S和L函数提供一个安全方式访问
	*/

	// 服务注册
	registerClient := consul.NewRegistryClient(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	serviceIdStr := fmt.Sprintf("%s", uuid.NewV4())
	if _err := registerClient.Register(
		global.ServerConfig.Host,
		global.ServerConfig.Port,
		global.ServerConfig.Name,
		global.ServerConfig.Tags,
		serviceIdStr,
	); _err != nil {
		zap.S().Panic("服务注册失败: ", _err.Error())
	}

	zap.S().Debugf("启动服务器, 端口: %d", global.ServerConfig.Port)
	go func() {
		if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
			zap.S().Panic("启动失败: ", err.Error())
		}
	}()
	// 接收服务终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if _err := registerClient.DeRegister(serviceIdStr); _err != nil {
		zap.S().Info("注销失败: ", _err.Error())
	} else {
		zap.S().Infof("%s -> 注销成功，id: %s", global.ServerConfig.Name, serviceIdStr)
	}
}
