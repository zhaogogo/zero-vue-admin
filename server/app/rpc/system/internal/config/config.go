package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	JWT struct {
		AccessSecret string
		AccessExpire int64
	}
	Mysql struct {
		DataSource string
	}
	CasbinConfig struct {
		Driver    string
		TableName string
		ModelPath string
	}
	CacheConf        cache.CacheConf
	RedisCacheConfig struct {
		Host string
		Pass string
	}
}
