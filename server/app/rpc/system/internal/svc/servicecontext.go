package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/model/system"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/config"
)

type ServiceContext struct {
	Config              config.Config
	Redis               *redis.Redis
	UserModel           system.UserModel
	UserRoleModel       system.UserRoleModel
	UserPageSetModel    system.UserPageSetModel
	UserMenuParamsModel system.UserMenuParamsModel
	RoleModel           system.RoleModel
	RoleMenuModel       system.RoleMenuModel
	MenuModel           system.MenuModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlxConn := sqlx.NewMysql(c.Mysql.DataSource)
	redis := redis.New(c.RedisCacheConfig.Host, redis.WithPass(c.RedisCacheConfig.Pass))
	return &ServiceContext{
		Config:              c,
		Redis:               redis,
		UserModel:           system.NewUserModel(sqlxConn, c.CacheConf),
		UserRoleModel:       system.NewUserRoleModel(sqlxConn, c.CacheConf),
		UserPageSetModel:    system.NewUserPageSetModel(sqlxConn, c.CacheConf),
		UserMenuParamsModel: system.NewUserMenuParamsModel(sqlxConn, c.CacheConf),
		RoleModel:           system.NewRoleModel(sqlxConn, c.CacheConf),
		RoleMenuModel:       system.NewRoleMenuModel(sqlxConn, c.CacheConf),
		MenuModel:           system.NewMenuModel(sqlxConn, c.CacheConf),
	}
}
