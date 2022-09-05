package sysuser

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysUserRoleModel = (*customSysUserRoleModel)(nil)

type (
	// SysUserRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysUserRoleModel.
	SysUserRoleModel interface {
		sysUserRoleModel
		FindByUserID(ctx context.Context, userId int64) ([]SysUserRole, error)
	}

	customSysUserRoleModel struct {
		*defaultSysUserRoleModel
	}
)

// NewSysUserRoleModel returns a model for the database table.
func NewSysUserRoleModel(conn sqlx.SqlConn, c cache.CacheConf) SysUserRoleModel {
	return &customSysUserRoleModel{
		defaultSysUserRoleModel: newSysUserRoleModel(conn, c),
	}
}

func (m *defaultSysUserRoleModel) FindByUserID(ctx context.Context, userId int64) ([]SysUserRole, error) {
	var resp []SysUserRole
	sql := fmt.Sprintf("SELECT %s FROM %s WHERE user_id = ?", sysUserRoleRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, sql, userId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
