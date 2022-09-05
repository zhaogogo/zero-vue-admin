package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/api/internal/config"
	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/api/internal/middleware"
	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/rpc/system/systemuser"
)

type ServiceContext struct {
	Config        config.Config
	CheckUrl      rest.Middleware
	Redis         *redis.Redis
	SystemUserRpc systemuser.SystemUser
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisClient := redis.New(c.Redis.Host, redisConfig(c))
	return &ServiceContext{
		Config:        c,
		CheckUrl:      middleware.NewCheckUrlMiddleware(redisClient).Handle,
		Redis:         redisClient,
		SystemUserRpc: systemuser.NewSystemUser(zrpc.MustNewClient(c.SystemUserRpcConf)),
	}
}

func redisConfig(c config.Config) redis.Option {
	return func(r *redis.Redis) {
		r.Type = redis.NodeType
		r.Pass = c.Redis.Pass
	}
}
