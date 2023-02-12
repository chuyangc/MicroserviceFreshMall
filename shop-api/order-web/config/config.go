package config

type GoodsSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

type OrderSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

type InventorySrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type AlipayConfig struct {
	AppID        string `mapstructure:"app_id" json:"app_id"`
	PrivateKey   string `mapstructure:"private_key" json:"private_key"`
	AliPublicKey string `mapstructure:"ali_public_key" json:"ali_public_key"`
	NotifyURL    string `mapstructure:"notify_url" json:"notify_url"`
	ReturnURL    string `mapstructure:"return_url" json:"return_url"`
}

type NacosConfig struct {
	Host      string `mapstructure:"host" json:"host"`
	Port      uint64 `mapstructure:"port" json:"port"`
	Namespace string `mapstructure:"namespace" json:"namespace"`
	User      string `mapstructure:"user" json:"user"`
	Password  string `mapstructure:"password" json:"password"`
	DataId    string `mapstructure:"dataid" json:"dataid"`
	Group     string `mapstructure:"group" json:"group"`
}

type ServerConfig struct {
	Name             string             `mapstructure:"name" json:"name"`
	Host             string             `mapstructure:"host" json:"host"`
	Port             int                `mapstructure:"port" json:"port"`
	Tags             []string           `mapstructure:"tags" json:"tags"`
	GoodsSrvInfo     GoodsSrvConfig     `mapstructure:"goods_srv" json:"goods_srv"`
	OrderSrvInfo     OrderSrvConfig     `mapstructure:"order_srv" json:"order_srv"`
	InventorySrvInfo InventorySrvConfig `mapstructure:"inventory_srv" json:"inventory_srv"`
	JWTInfo          JWTConfig          `mapstructure:"jwt" json:"jwt"`
	ConsulInfo       ConsulConfig       `mapstructure:"consul" json:"consul"`
	NacosInfo        NacosConfig        `mapstructure:"nacos" json:"nacos"`
	AliPayInfo       AlipayConfig       `mapstructure:"alipay" json:"alipay"`
}
