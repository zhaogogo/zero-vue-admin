package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/rpc/model/systemusermodel"
	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/rpc/system/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserModel systemusermodel.SystemUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: systemusermodel.NewSystemUserModel(sqlConn),
	}
}
