package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/usercenter/api/internal/config"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/usercenter/model"
)

type ServiceContext struct {
	Config              config.Config
	RedisClient         *redis.Redis
	UserModel           model.UserModel
	UserPageSetModel    model.UserPageSetModel
	UserMenuParamsModel model.UserMenuParamsModel
	RoleModel           model.RoleModel
	UserRoleModel       model.UserRoleModel
	RoleMenuModel       model.RoleMenuModel
	MenuModel           model.MenuModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisClient := redis.New(c.Redis.Host, redisOpts(c))
	sqlxConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:      c,
		RedisClient: redisClient,

		UserModel:           model.NewUserModel(sqlxConn, c.CacheConf),
		UserPageSetModel:    model.NewUserPageSetModel(sqlxConn, c.CacheConf),
		UserMenuParamsModel: model.NewUserMenuParamsModel(sqlxConn, c.CacheConf),
		RoleModel:           model.NewRoleModel(sqlxConn, c.CacheConf),
		UserRoleModel:       model.NewUserRoleModel(sqlxConn, c.CacheConf),
		RoleMenuModel:       model.NewRoleMenuModel(sqlxConn, c.CacheConf),
		MenuModel:           model.NewMenuModel(sqlxConn, c.CacheConf),
	}
}

func redisOpts(c config.Config) redis.Option {
	return func(r *redis.Redis) {
		r.Type = redis.NodeType
		r.Pass = c.Redis.Pass
	}
}
