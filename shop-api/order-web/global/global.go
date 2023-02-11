package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"shop-api/order-web/config"
	"shop-api/order-web/proto"
)

var (
	ServerConfig        *config.ServerConfig = &config.ServerConfig{}
	Trans               ut.Translator
	GoodsSrvClient      proto.GoodsClient
	OrderSrvClient      proto.OrderClient
	NacosConfig         *config.NacosConfig = &config.NacosConfig{}
	GoodsNacosSrvConfig naming_client.INamingClient
	InventorySrvClient  proto.InventoryClient
)
