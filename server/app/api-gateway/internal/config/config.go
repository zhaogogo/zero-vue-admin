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
	MonitoringConfig MonitoringConfig

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

type MonitoringConfig struct {
	AlertmanagerSlienceURL AlertmanagerSlienceURL
	NotifyURL              string
	AggregationNotify      string
	AggregationSeverity    string
	NotifyTemplatePath     string
}

type AlertmanagerSlienceURL struct {
	ZW string
	YZ string
}
