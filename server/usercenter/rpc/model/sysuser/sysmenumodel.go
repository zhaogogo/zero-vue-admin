package sysuser

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ SysMenuModel = (*customSysMenuModel)(nil)

type (
	// SysMenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysMenuModel.
	SysMenuModel interface {
		sysMenuModel
		FindMenusByRoles(ctx context.Context, roleIds ...string) ([]*SysMenu, error)
	}

	customSysMenuModel struct {
		*defaultSysMenuModel
	}
)

// NewSysMenuModel returns a model for the database table.
func NewSysMenuModel(conn sqlx.SqlConn, c cache.CacheConf) SysMenuModel {
	return &customSysMenuModel{
		defaultSysMenuModel: newSysMenuModel(conn, c),
	}
}

func (m *defaultSysMenuModel) FindMenusByRoles(ctx context.Context, roleIds ...string) ([]*SysMenu, error) {
	sql := fmt.Sprintf("SELECT %s FROM %s WHERE id IN (SELECT menu_id FROM sys_role_menu WHERE role_id IN (?)) AND delete_at IS NULL", sysMenuRows, m.table)
	var resp []*SysMenu
	err := m.QueryRowsNoCacheCtx(ctx, &resp, sql, strings.Join(roleIds, ","))
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
