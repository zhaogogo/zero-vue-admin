package system

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var (
	cacheChaosSystemRoleMenuRoleIdPrefix = "cache:chaosSystem:roleMenu:role_id:"
)

var _ RoleMenuModel = (*customRoleMenuModel)(nil)

type (
	// RoleMenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRoleMenuModel.
	RoleMenuModel interface {
		roleMenuModel
		FindByRoleID(ctx context.Context, roleID uint64) ([]RoleMenu, error)
	}

	customRoleMenuModel struct {
		*defaultRoleMenuModel
	}
)

// NewRoleMenuModel returns a model for the database table.
func NewRoleMenuModel(conn sqlx.SqlConn, c cache.CacheConf) RoleMenuModel {
	return &customRoleMenuModel{
		defaultRoleMenuModel: newRoleMenuModel(conn, c),
	}
}

func (m *defaultRoleMenuModel) FindByRoleID(ctx context.Context, roleID uint64) ([]RoleMenu, error) {
	chaosSystemRoleMenuRuleIdKey := fmt.Sprintf("%s%v", cacheChaosSystemRoleMenuRoleIdPrefix, roleID)
	var resp []RoleMenu
	query := fmt.Sprintf("SELECT %s FROM %s where `role_id` = ?", roleMenuRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, roleID)
	if err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return nil, ErrNotFound
	}
	m.SetCacheCtx(ctx, chaosSystemRoleMenuRuleIdKey, resp)
	return resp, nil
}
