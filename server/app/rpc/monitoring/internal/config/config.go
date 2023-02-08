package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	CacheConf        cache.CacheConf
	RedisCacheConfig struct {
		Host string
		Pass string
	}
	Mysql struct {
		DataSource string
	}
}
