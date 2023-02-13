package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

// TODO nacos 注册中心工具 有待完善

var namingclient *naming_client.NamingClient

type Registry struct {
	Host string
	Port int
}

type RegistryClient interface {
	Register(address string, port int, name string, tags []string, id string) error
	DeRegister(serviceId string) error
}

func (r *Registry) Register(address string, port int, name string, tags []string, id string) error {
	// 创建ServerConfig
	SC := []constant.ServerConfig{
		*constant.NewServerConfig(address, uint64(port), constant.WithContextPath("/nacos")),
	}

	CC := *constant.NewClientConfig(
		constant.WithNamespaceId(id),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)

	namingclient, _ := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &CC,
			ServerConfigs: SC,
		},
	)

	_, err := namingclient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          r.Host,
		Port:        8848,
		ServiceName: name,
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		//Ephemeral:   true,
		//Metadata:    map[string]string{"idc":"shanghai"},
		//ClusterName: "cluster-a", // 默认值DEFAULT
		//GroupName:   "group-a",   // 默认值DEFAULT_GROUP
	})

	return err
}

func (r *Registry) DeRegister(serviceId string) error {
	_, err := namingclient.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          r.Host,
		Port:        8848,
		ServiceName: serviceId,
		//Ephemeral:   true,
		//Cluster:     "cluster-a", // 默认值DEFAULT
		//GroupName:   "group-a",   // 默认值DEFAULT_GROUP
	})
	return err
}
