package svc

import (
	casbinmysqladapter "github.com/Blank-Xu/sql-adapter"
	"github.com/casbin/casbin/v2"
	casbinmodel "github.com/casbin/casbin/v2/model"
	casbinutil "github.com/casbin/casbin/v2/util"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/config"
	model "github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/model"
	"strings"
)

type ServiceContext struct {
	Config config.Config
	Redis  *redis.Redis

	SyncedEnforcer      *casbin.SyncedEnforcer
	CasbinRuleModel     model.CasbinRuleModel
	UserModel           model.UserModel
	UserRoleModel       model.UserRoleModel
	UserPageSetModel    model.UserPageSetModel
	UserMenuParamsModel model.UserMenuParamsModel
	RoleModel           model.RoleModel
	RoleMenuModel       model.RoleMenuModel
	MenuModel           model.MenuModel
	APIModel            model.ApiModel
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
		CasbinRuleModel:     model.NewCasbinRuleModel(sqlxConn),
		UserModel:           model.NewUserModel(sqlxConn, c.CacheConf),
		UserRoleModel:       model.NewUserRoleModel(sqlxConn, c.CacheConf),
		UserPageSetModel:    model.NewUserPageSetModel(sqlxConn, c.CacheConf),
		UserMenuParamsModel: model.NewUserMenuParamsModel(sqlxConn, c.CacheConf),
		RoleModel:           model.NewRoleModel(sqlxConn, c.CacheConf),
		RoleMenuModel:       model.NewRoleMenuModel(sqlxConn, c.CacheConf),
		MenuModel:           model.NewMenuModel(sqlxConn, c.CacheConf),
		APIModel:            model.NewApiModel(sqlxConn, c.CacheConf),
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
	m, err := casbinmodel.NewModelFromFile(c.CasbinConfig.ModelPath)
	//m, err := casbinmodel.NewModelFromString()
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
