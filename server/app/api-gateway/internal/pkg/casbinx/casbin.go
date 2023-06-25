package casbinx

import (
	"database/sql"
	casbinmysqladapter "github.com/Blank-Xu/sql-adapter"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"strings"
	"sync"
	"time"
)

var Casbin = &Casbinx{}

type Casbinx struct {
	table      string
	db         *sql.DB
	driverName string
	m          model.Model
	a          *casbinmysqladapter.Adapter
	enforcer   *casbin.Enforcer
	once       sync.Once
	mu         sync.RWMutex
}

func (casbinx *Casbinx) LoadPolicy() error {
	casbinx.mu.Lock()
	defer casbinx.mu.Unlock()
	return casbinx.enforcer.LoadPolicy()
}

func (casbinx *Casbinx) Enforcex(rvals ...interface{}) (bool, error) {
	casbinx.mu.RLock()
	defer casbinx.mu.RUnlock()
	return casbinx.enforcer.Enforce(rvals...)
}

func (casbinx *Casbinx) Enforce(rvals ...interface{}) (bool, error) {
	var err error
	casbinx.once.Do(func() {
		casbinx.m, err = model.NewModelFromString(modelString)
		if err != nil {
			err = errors.Wrapf(err, "casbin创建model失败")
			return
		}

		casbinx.a, err = casbinmysqladapter.NewAdapter(casbinx.db, casbinx.driverName, casbinx.table)
		if err != nil {
			err = errors.Wrapf(err, "casbin mysql adapter创建失败")
			return
		}
		casbin.NewSyncedEnforcer()
		casbinx.enforcer, err = casbin.NewEnforcer(casbinx.m, casbinx.a)
		if err != nil {
			err = errors.Wrapf(err, "创建syncEnforcer失败")
			return
		}

		casbinx.enforcer.AddFunction("ParamsMatch", ParamsMatchFunc)
	})
	if err != nil {
		return false, err
	}
	err = casbinx.LoadPolicy()
	if err != nil {
		return false, errors.Wrap(err, "加载Policy失败")
	}
	casbinx.mu.RLock()
	defer casbinx.mu.RUnlock()
	success, err := casbinx.enforcer.Enforce(rvals...)
	if err != nil {
		return false, errors.Wrap(err, "权限验证错误")
	}

	return success, nil
}
func (casbinx *Casbinx) SetUp2(sql *sql.DB) error {
	casbinx.db = sql
	casbinx.driverName = "mysql"
	casbinx.table = "casbin_rule"
	return nil
}
func (casbinx *Casbinx) SetUp(dsn string) error {
	sql, err := sql.Open("mysql", dsn)
	if err != nil {
		return errors.Wrapf(err, "casbin sql conn pool创建失败")
	}
	if err := sql.Ping(); err != nil {
		return errors.Wrapf(err, "casbin sql conn无法连接数据库")
	}
	sql.SetMaxIdleConns(64)
	sql.SetMaxOpenConns(64)
	sql.SetConnMaxLifetime(time.Minute)

	casbinx.db = sql
	casbinx.driverName = "mysql"
	casbinx.table = "casbin_rule"
	return nil
}

func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return ParamsMatch(name1, name2), nil
}

func ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	// 剥离路径后再使用casbin的keyMatch2
	return util.KeyMatch2(key1, key2)
}

func Enforce(rvals ...interface{}) (bool, error) {
	return Casbin.Enforce(rvals...)
}
