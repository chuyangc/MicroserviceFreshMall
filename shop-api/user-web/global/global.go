package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"shop-api/user-web/config"
	"shop-api/user-web/proto"
)

var (
	ServerConfig       *config.ServerConfig = &config.ServerConfig{}
	Trans              ut.Translator
	UserSrvClient      proto.UserClient
	NacosConfig        *config.NacosConfig = &config.NacosConfig{}
	UserNacosSrvConfig naming_client.INamingClient
)
