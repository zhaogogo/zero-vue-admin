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
		FindByMenuID_NC(ctx context.Context, menuid uint64) ([]RoleMenu, error)
		FindByRoleID(ctx context.Context, redis *redis.Redis, roleID uint64) ([]RoleMenu, error)

		TransDeleteByMenuId(ctx context.Context, session sqlx.Session, menuid uint64) error
		TransDeleteByRoleId(ctx context.Context, session sqlx.Session, roleid uint64) error
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

func (m *defaultRoleMenuModel) FindByMenuID_NC(ctx context.Context, menuid uint64) ([]RoleMenu, error) {
	var resp []RoleMenu
	query := fmt.Sprintf("SELECT %s FROM %s WHERE `menu_id` = ?", roleMenuRows, m.table)

	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, menuid)
	if err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return nil, ErrNotFound
	}

	return resp, nil
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

func (m *defaultRoleMenuModel) TransDeleteByMenuId(ctx context.Context, session sqlx.Session, menuid uint64) error {
	cacheChaosSystemRoleMenu_RoleId_Keys := []string{}
	roleidSet := make(map[uint64]int64)

	var rolemenu []RoleMenu
	query := fmt.Sprintf("SELECT %s FROM %s WHERE `menu_id` = ?", roleMenuRows, m.table)
	err := session.QueryRowsCtx(ctx, &rolemenu, query, menuid)
	if err != nil {
		return err
	}
	for _, rolemenu := range rolemenu {
		roleidSet[rolemenu.RoleId] += 1
	}
	for roleid, _ := range roleidSet {
		cacheChaosSystemRoleMenu_RoleId_Keys = append(cacheChaosSystemRoleMenu_RoleId_Keys, fmt.Sprintf("%s%v", cacheChaosSystemRoleMenuRoleIdPrefix, roleid))
	}

	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `menu_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, menuid)
	}, cacheChaosSystemRoleMenu_RoleId_Keys...)
	return err
}

func (m *defaultRoleMenuModel) TransDeleteByRoleId(ctx context.Context, session sqlx.Session, roleid uint64) error {
	cacheChaosSystemRoleMenu_RoleId_Keys := fmt.Sprintf("%s%v", cacheChaosSystemRoleMenuRoleIdPrefix, roleid)

	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `role_id` = ?", m.table)
		return session.ExecCtx(ctx, query, roleid)
	}, cacheChaosSystemRoleMenu_RoleId_Keys)
	return err
}
