package initialize

import (
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"shop-api/goods-web/global"
)

//初始化配置

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	debug := GetEnvInfo("SHOP_DEBUG")
	configFileNamePrefix := "config"
	configFileName := fmt.Sprintf("goods-web/%s-debug.yaml", configFileNamePrefix)
	if debug {
		configFileName = fmt.Sprintf("goods-web/%s-pro.yaml", configFileNamePrefix)
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

func InitNacosConfig() {
	debug := GetEnvInfo("SHOP_DEBUG")
	configFileNamePrefix := "config"
	configFileName := fmt.Sprintf("goods-web/%s-debug.yaml", configFileNamePrefix)
	if debug {
		configFileName = fmt.Sprintf("goods-web/%s-pro.yaml", configFileNamePrefix)
	}

	v := viper.New()
	v.SetConfigFile(configFileName) //设置读取的文件名
	v.SetConfigType("yaml")         //设置文件的类型
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	// 使用全局变量 - ServerConfig
	if err := v.Unmarshal(global.NacosConfig); err != nil {
		panic(err)
	}
	zap.S().Infof("配置信息加载成功 -> %v", global.NacosConfig)

	// 从Nacos中读取配置信息
	// 创建serverConfig
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(
			global.NacosConfig.Host,
			global.NacosConfig.Port,
		),
	}

	// 创建clientConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(global.NacosConfig.Namespace), //当namespace是public时，此处填空字符串。
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)

	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic(err)
	}

	content, _err := configClient.GetConfig(vo.ConfigParam{
		DataId: global.NacosConfig.DataId,
		Group:  global.NacosConfig.Group})
	if _err != nil {
		panic(_err)
	}

	__err := json.Unmarshal([]byte(content), &global.ServerConfig)
	if __err != nil {
		zap.S().Fatalf("读取Nacos配置失败 -> %s", err.Error())
	}
	//fmt.Println(&global.ServerConfig)
}
