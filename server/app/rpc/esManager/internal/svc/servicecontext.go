package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/esManager/internal/config"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/esManager/model"
)

type ServiceContext struct {
	Config      config.Config
	ESConnModel model.EsConnModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlxConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:      c,
		ESConnModel: model.NewEsConnModel(sqlxConn, c.CacheConf),
	}
}
