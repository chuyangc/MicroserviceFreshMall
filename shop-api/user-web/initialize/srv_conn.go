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

	"shop-api/user-web/global"
	"shop-api/user-web/proto"
)

// 初始化服务发现

func InitConsulLBSrvConn() {
	consulInfo := global.ServerConfig.ConsulInfo
	userConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.UserSrvInfo.Name),
		/*
			The function insecure.NewCredentials returns an implementation of credentials.TransportCredentials.
			use it as a DialOption with grpc.WithTransportCredentials.
			Deprecated: use WithTransportCredentials and insecure.NewCredentials() to instead of it
			grpc.Dial(":port", grpc.WithTransportCredentials(insecure.NewCredentials()))
		*/
		//grpc.WithInsecure(), //弃用的方法
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接 [用户服务失败]")
	}

	userSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = userSrvClient
}

func InitNacosLBSrvConn() {
	// TODO 有待完善的获取Nacos用户服务,且未完成测试
	// 初始化UserNacosSrvConfig实例
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
	global.UserNacosSrvConfig = client
	//使用初始化完成的UserNacosSrvConfig实例获取服务
	instance, _err := global.UserNacosSrvConfig.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: global.ServerConfig.Name,
		//Clusters:    []string{"DEMO_SERVER"},
		//GroupName:   "DEMO_SERVER_GROUP",
	})
	if _err != nil {
		zap.S().Error(err)
	}
	// Nacos2.0版本支持gRPC
	userConn, __err := grpc.Dial(
		fmt.Sprintf("%s:%d", instance.Ip, instance.Port),
		//grpc.WithInsecure(), //弃用的方法
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`))
	if __err != nil {
		zap.S().Fatal("[InitSrvConn] 连接 [用户服务失败]")
	}
	userSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = userSrvClient
}

func InitConsulSrvConn() {
	//从注册中心获取到用户服务的信息
	cfg := api.DefaultConfig()
	consulInfo := global.ServerConfig.ConsulInfo
	cfg.Address = fmt.Sprintf("%s:%d", consulInfo.Host, consulInfo.Port)

	userSrvHost := ""
	userSrvPort := 0
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf("Service == \"%s\"", global.ServerConfig.UserSrvInfo.Name))
	//data, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, global.ServerConfig.UserSrvInfo.Name))
	if err != nil {
		panic(err)
	}
	for _, value := range data {
		userSrvHost = value.Address
		userSrvPort = value.Port
		break
	}
	if userSrvHost == "" {
		zap.S().Fatal("[InitSrvConn] 连接 [用户服务失败]")
		return
	}

	//拨号连接用户grpc服务器
	userConn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", userSrvHost, userSrvPort),
		//grpc.WithInsecure(), //弃用的方法
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		zap.S().Errorw("[GetUserList] 连接 [用户服务失败]",
			"msg", err.Error(),
		)
	}
	//1. 后续的用户服务下线了 2. 改端口了 3. 改ip了 -> 负载均衡
	//2. 已经事先创立好了连接，不用进行再次tcp的三次握手
	//3. 一个连接多个协程共用，存在性能问题 - 连接池
	userSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = userSrvClient
}
