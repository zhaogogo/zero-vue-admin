package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/monitoring/internal/config"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/monitoring/model"
)

type ServiceContext struct {
	Config         config.Config
	Redis          *redis.Redis
	AlertRuleModel model.AlertRuleModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlxConn := sqlx.NewMysql(c.Mysql.DataSource)
	redisConn := redis.New(c.RedisCacheConfig.Host, redis.WithPass(c.RedisCacheConfig.Pass))

	return &ServiceContext{
		Config:         c,
		Redis:          redisConn,
		AlertRuleModel: model.NewAlertRuleModel(sqlxConn, c.CacheConf),
	}
}
