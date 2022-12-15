package system

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
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
		FindByRoleID(ctx context.Context, redis *redis.Redis, roleID uint64) ([]RoleMenu, error)
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

func (m *defaultRoleMenuModel) FindByRoleID(ctx context.Context, redis *redis.Redis, roleID uint64) ([]RoleMenu, error) {
	chaosSystemRoleMenuRuleIdKey := fmt.Sprintf("%s%v", cacheChaosSystemRoleMenuRoleIdPrefix, roleID)
	var resp []RoleMenu
	err := m.GetCacheCtx(ctx, chaosSystemRoleMenuRuleIdKey, &resp)
	if err == nil {
		return resp, nil
	}
	if err.Error() == "placeholder" {
		return nil, ErrNotFound
	}
	if err == sql.ErrNoRows {
		query := fmt.Sprintf("SELECT %s FROM %s where `role_id` = ?", roleMenuRows, m.table)
		err := m.QueryRowsNoCacheCtx(ctx, &resp, query, roleID)
		if err != nil {
			return nil, err
		}
		if len(resp) == 0 {
			err := redis.Setex(chaosSystemRoleMenuRuleIdKey, "*", int(unstable.AroundDuration(cacheOption.NotFoundExpiry).Seconds()))
			if err != nil {
				logx.Errorf("设置缓存失败, key: %v, error: %v", chaosSystemRoleMenuRuleIdKey, err)
			}
			return nil, ErrNotFound
		}
		err = m.SetCacheCtx(ctx, chaosSystemRoleMenuRuleIdKey, resp)
		if err != nil {
			logx.Errorf("设置缓存失败, key: %v, error: %v", chaosSystemRoleMenuRuleIdKey, err)
		}
		return resp, nil
	}

	return nil, err
}
