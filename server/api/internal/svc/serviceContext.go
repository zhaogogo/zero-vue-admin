package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zhaoqiang0201/zero-vue-admin/server/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Auth   struct {
		AccessSecret string
		AccessExpire int64
	}
	Redis *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisClient := redis.New(c.Redis.Host, redisConfig(c))
	return &ServiceContext{
		Config: c,
		Redis:  redisClient,
	}
}

func redisConfig(c config.Config) redis.Option {
	return func(r *redis.Redis) {
		r.Type = redis.NodeType
		r.Pass = c.Redis.Pass
	}
}
