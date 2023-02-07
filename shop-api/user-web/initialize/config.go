package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"shop-api/user-web/global"
)

//初始化配置
func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	debug := GetEnvInfo("SHOP_DEBUG")
	configFileNamePrefix := "config"
	configFileName := fmt.Sprintf("user-web/%s-debug.yaml", configFileNamePrefix)
	if debug {
		configFileName = fmt.Sprintf("user-web/%s-pro.yaml", configFileNamePrefix)
	}

	v := viper.New()
	v.SetConfigFile(configFileName) //设置读取的文件名
	v.SetConfigType("yaml")         //设置文件的类型
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	// 使用全局变量 - ServerConfig
	if err := v.Unmarshal(global.ServerConfig); err != nil {
		panic(err)
	}
	zap.S().Infof("配置信息: %v", global.ServerConfig)
	// viper的功能 - 动态监控变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		zap.S().Infof("配置信息产生变化: %s", e.Name)
		_ = v.ReadInConfig() // 读取配置数据
		_ = v.Unmarshal(global.ServerConfig)
		zap.S().Infof("配置信息: %v", global.ServerConfig)
	})
}
