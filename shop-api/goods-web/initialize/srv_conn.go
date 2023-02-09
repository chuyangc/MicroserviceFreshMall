package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"shop-api/goods-web/global"
	"shop-api/goods-web/proto"
)

// 初始化服务发现

func InitConsulLBSrvConn() {
	consulInfo := global.ServerConfig.ConsulInfo
	goodsConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.GoodsSrvInfo.Name),
		//grpc.WithInsecure(), //弃用的方法
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接 [商品服务失败]")
	}

	goodsSrvClient := proto.NewGoodsClient(goodsConn)
	global.GoodsSrvClient = goodsSrvClient
}

func InitNacosLBSrvConn() {
	// TODO 有待完善的获取Nacos商品服务,且未完成测试
	// 初始化GoodsNacosSrvConfig实例
	// 创建ServerConfig
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(global.ServerConfig.NacosInfo.Host, global.ServerConfig.NacosInfo.Port, constant.WithContextPath("/nacos")),
	}

	// TODO 创建ClientConfig，参数可加入配置文件中，有待完善
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId("838b806f-20e6-4c3b-abf2-0e3b3bc73dcc"),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)

	// 创建naming client
	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		zap.S().Fatal("获取服务失败")
		//panic(err)
	}
	global.GoodsNacosSrvConfig = client
	//使用初始化完成的GoodsNacosSrvConfig实例获取服务
	instance, _err := global.GoodsNacosSrvConfig.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: global.ServerConfig.Name,
		//Clusters:    []string{"DEMO_SERVER"},
		//GroupName:   "DEMO_SERVER_GROUP",
	})
	if _err != nil {
		zap.S().Error(err)
	}
	// Nacos2.0版本支持gRPC
	goodsConn, __err := grpc.Dial(
		fmt.Sprintf("%s:%d", instance.Ip, instance.Port),
		//grpc.WithInsecure(), //弃用的方法
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`))
	if __err != nil {
		zap.S().Fatal("[InitSrvConn] 连接 [商品服务失败]")
	}
	goodsSrvClient := proto.NewGoodsClient(goodsConn)
	global.GoodsSrvClient = goodsSrvClient
}

func InitConsulSrvConn() {
	//从注册中心获取到商品服务的信息
	cfg := api.DefaultConfig()
	consulInfo := global.ServerConfig.ConsulInfo
	cfg.Address = fmt.Sprintf("%s:%d", consulInfo.Host, consulInfo.Port)

	GoodsSrvHost := ""
	GoodsSrvPort := 0
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf("Service == \"%s\"", global.ServerConfig.GoodsSrvInfo.Name))
	//data, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, global.ServerConfig.GoodsSrvInfo.Name))
	if err != nil {
		panic(err)
	}
	for _, value := range data {
		GoodsSrvHost = value.Address
		GoodsSrvPort = value.Port
		break
	}
	if GoodsSrvHost == "" {
		zap.S().Fatal("[InitSrvConn] 连接 [商品服务失败]")
		return
	}

	//拨号连接商品grpc服务器
	goodsConn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", GoodsSrvHost, GoodsSrvPort),
		//grpc.WithInsecure(), //弃用的方法
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		zap.S().Errorw("[GetGoodsList] 连接 [商品服务失败]",
			"msg", err.Error(),
		)
	}
	goodsSrvClient := proto.NewGoodsClient(goodsConn)
	global.GoodsSrvClient = goodsSrvClient
}
