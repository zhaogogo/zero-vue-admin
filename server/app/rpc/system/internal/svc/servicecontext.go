package svc

import (
	casbinmysqladapter "github.com/Blank-Xu/sql-adapter"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	casbinutil "github.com/casbin/casbin/v2/util"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/config"
	model2 "github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/model"
	"strings"
)

type ServiceContext struct {
	Config config.Config
	Redis  *redis.Redis

	SyncedEnforcer      *casbin.SyncedEnforcer
	CasbinRuleModel     model2.CasbinRuleModel
	UserModel           model2.UserModel
	UserRoleModel       model2.UserRoleModel
	UserPageSetModel    model2.UserPageSetModel
	UserMenuParamsModel model2.UserMenuParamsModel
	RoleModel           model2.RoleModel
	RoleMenuModel       model2.RoleMenuModel
	MenuModel           model2.MenuModel
	APIModel            model2.ApiModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlxConn := sqlx.NewMysql(c.Mysql.DataSource)
	redisConn := redis.New(c.RedisCacheConfig.Host, redis.WithPass(c.RedisCacheConfig.Pass))
	syncEnforcer, err := NewCasbinSyncedEnforcer(c, sqlxConn)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:              c,
		Redis:               redisConn,
		SyncedEnforcer:      syncEnforcer,
		CasbinRuleModel:     model2.NewCasbinRuleModel(sqlxConn),
		UserModel:           model2.NewUserModel(sqlxConn, c.CacheConf),
		UserRoleModel:       model2.NewUserRoleModel(sqlxConn, c.CacheConf),
		UserPageSetModel:    model2.NewUserPageSetModel(sqlxConn, c.CacheConf),
		UserMenuParamsModel: model2.NewUserMenuParamsModel(sqlxConn, c.CacheConf),
		RoleModel:           model2.NewRoleModel(sqlxConn, c.CacheConf),
		RoleMenuModel:       model2.NewRoleMenuModel(sqlxConn, c.CacheConf),
		MenuModel:           model2.NewMenuModel(sqlxConn, c.CacheConf),
		APIModel:            model2.NewApiModel(sqlxConn, c.CacheConf),
	}
}

func NewCasbinSyncedEnforcer(c config.Config, sqlxConn sqlx.SqlConn) (*casbin.SyncedEnforcer, error) {
	sqlConn, err := sqlxConn.RawDB()
	if err != nil {
		return nil, err
	}
	if c.CasbinConfig.Driver != "mysql" {
		return nil, errors.New("仅支持mysql驱动")
	}
	m, err := model.NewModelFromString(c.CasbinConfig.Model)
	if err != nil {
		return nil, errors.Wrap(err, "casbin创建model失败")
	}
	a, err := casbinmysqladapter.NewAdapter(sqlConn, c.CasbinConfig.Driver, c.CasbinConfig.TableName)
	if err != nil {
		return nil, errors.Wrapf(err, "casbin_mysql_adapter创建失败")
	}

	syncEnforcer, err := casbin.NewSyncedEnforcer(m, a)
	if err != nil {
		return nil, errors.Wrapf(err, "创建syncEnforcer失败")
	}

	syncEnforcer.AddFunction("ParamsMatch", ParamsMatchFunc)
	_ = syncEnforcer.LoadPolicy()
	return syncEnforcer, err
}

func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return ParamsMatch(name1, name2), nil
}

func ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	// 剥离路径后再使用casbin的keyMatch2
	return casbinutil.KeyMatch2(key1, key2)
}
