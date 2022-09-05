package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/rpc/model/sysuser"
	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/rpc/system/internal/config"
)

type ServiceContext struct {
	Config config.Config

	UserModel     sysuser.SysUserModel
	UserRoleModel sysuser.SysUserRoleModel
	RoleModel     sysuser.SysRoleModel
	RoleMenuModel sysuser.SysRoleMenuModel
	MenuModel     sysuser.SysMenuModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:        c,
		UserModel:     sysuser.NewSysUserModel(sqlConn, c.RedisCache),
		UserRoleModel: sysuser.NewSysUserRoleModel(sqlConn, c.RedisCache),
		RoleModel:     sysuser.NewSysRoleModel(sqlConn, c.RedisCache),
		RoleMenuModel: sysuser.NewSysRoleMenuModel(sqlConn),
		MenuModel:     sysuser.NewSysMenuModel(sqlConn, c.RedisCache),
	}
}
