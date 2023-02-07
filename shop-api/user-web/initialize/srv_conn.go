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

	"shop-api/user-web/global"
	"shop-api/user-web/proto"
)

// 初始化服务注册与发现

func InitLBSrvConn() {
	consulInfo := global.ServerConfig.ConsulInfo
	userConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.UserSrvInfo.Name),
		grpc.WithInsecure(),
		//insecure.NewCredentials(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接 [用户服务失败]")
	}

	userSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = userSrvClient
}

func InitNacosSrvConn() {
	GetNacosSrvConn()
	//拨号连接用户grpc服务器
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d/%s", global.ServerConfig.UserSrvInfo.Host,
		global.ServerConfig.UserSrvInfo.Port, global.ServerConfig.UserSrvInfo.Name), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList] 连接 [用户服务失败]",
			"msg", err.Error())
	}
	//生成grpc的client并调用接口
	userSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = userSrvClient
}

func GetNacosServices(srvname string) {
	_, err := global.UserNacosSrvConfig.GetService(vo.GetServiceParam{
		ServiceName: "demo.go",
		Clusters:    []string{"cluster-a"}, // 默认值DEFAULT
		GroupName:   "group-a",             // 默认值DEFAULT_GROUP
	})
	if err != nil {
		zap.S().Fatal("获取Nacos服务列表失败")
	}
}

func GetNacosSrvConn() {
	//create ServerConfig
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("127.0.0.1", 8848, constant.WithContextPath("/nacos")),
	}

	//create ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId("838b806f-20e6-4c3b-abf2-0e3b3bc73dcc"),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)

	// create naming client
	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	global.UserNacosSrvConfig = client

	if err != nil {
		zap.S().Fatal("[InitNacosSrvConn] 连接 [用户服务失败]")
		//panic(err)
	}
	success, _err := client.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "127.0.0.1",
		Port:        8021,
		ServiceName: "user-web",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		//Ephemeral:   true,
		//Metadata:    map[string]string{"idc": "shanghai"},
		//ClusterName: "cluster-a", // 默认值DEFAULT
		//GroupName:   "group-a",   // 默认值DEFAULT_GROUP
	})
	if !success {
		zap.S().Fatal(_err)
	}
	zap.S().Info("user-web 服务注册 成功")
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
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", userSrvHost, userSrvPort), grpc.WithInsecure())
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
