package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	Mysql Mysql
	Auth  struct {
		AccessSecret string
		AccessExpire int64
	}

	SystemAdminRpcConf zrpc.RpcClientConf
	ESManagerRpcConf   zrpc.RpcClientConf
	MonitoringRpcConf  zrpc.RpcClientConf
}

type Mysql struct {
	System     System
	Monitoring Monitoring
	LogPath    string
}
type System struct {
	DataSource string
}

type Monitoring struct {
	DataSource string
}
