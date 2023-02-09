package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"shop-api/goods-web/config"
	"shop-api/goods-web/proto"
)

var (
	ServerConfig        *config.ServerConfig = &config.ServerConfig{}
	Trans               ut.Translator
	GoodsSrvClient      proto.GoodsClient
	NacosConfig         *config.NacosConfig = &config.NacosConfig{}
	GoodsNacosSrvConfig naming_client.INamingClient
)
