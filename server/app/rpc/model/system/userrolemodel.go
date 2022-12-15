package system

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var (
	cacheChaosSystemUserRoleUserIdPrefix = "cache:chaosSystem:userRole:user_id:"
)

var _ UserRoleModel = (*customUserRoleModel)(nil)

type (
	// UserRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserRoleModel.
	UserRoleModel interface {
		userRoleModel
		FindByUserID(ctx context.Context, redis *redis.Redis, userID uint64) ([]UserRole, error)
		TransCtx(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error
		TransDeleteByUserID(ctx context.Context, session sqlx.Session, userID uint64) error
		TranInsertUserIDRoleIDs(ctx context.Context, session sqlx.Session, userID uint64, role_id_s []uint64) (sql.Result, error)
	}

	customUserRoleModel struct {
		*defaultUserRoleModel
	}
)

// NewUserRoleModel returns a model for the database table.
func NewUserRoleModel(conn sqlx.SqlConn, c cache.CacheConf) UserRoleModel {
	return &customUserRoleModel{
		defaultUserRoleModel: newUserRoleModel(conn, c),
	}
}

func (m *defaultUserRoleModel) TransCtx(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *defaultUserRoleModel) TransDeleteByUserID(ctx context.Context, session sqlx.Session, userID uint64) error {
	userChaosSystemUserRoleUserIDKey := fmt.Sprintf("%s%v", cacheChaosSystemUserRoleUserIdPrefix, userID)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		query := fmt.Sprintf("DELETE FROM %s WHERE `user_id` = ?", m.table)
		return session.ExecCtx(ctx, query, userID)
	}, userChaosSystemUserRoleUserIDKey)

	return err
}

func (m *defaultUserRoleModel) TranInsertUserIDRoleIDs(ctx context.Context, session sqlx.Session, userID uint64, role_id_s []uint64) (sql.Result, error) {
	userChaosSystemUserRoleUserIDKey := fmt.Sprintf("%s%v", cacheChaosSystemUserRoleUserIdPrefix, userID)
	query := fmt.Sprintf("insert into %s (%s) values ", m.table, userRoleRowsExpectAutoSet)
	values := []string{}
	for _, role_id := range role_id_s {
		values = append(values, fmt.Sprintf("(%v, %v)", userID, role_id))
	}
	query += strings.Join(values, ", ")

	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		return session.ExecCtx(ctx, query)
	}, userChaosSystemUserRoleUserIDKey)
	return ret, err
}

func (m *defaultUserRoleModel) FindByUserID(ctx context.Context, redis *redis.Redis, userID uint64) ([]UserRole, error) {
	var resp []UserRole
	chaosSystemUserRoleUserIDKey := fmt.Sprintf("%s%v", cacheChaosSystemUserRoleUserIdPrefix, userID)
	err := m.GetCacheCtx(ctx, chaosSystemUserRoleUserIDKey, &resp)
	if err == nil {
		return resp, nil //查询成功
	}
	if err.Error() == "placeholder" {
		return nil, ErrNotFound
	}
	if err == sql.ErrNoRows { //redis查询
		query := fmt.Sprintf("SELECT %s FROM %s WHERE `user_id` = ?", userRoleRows, m.table)
		err = m.QueryRowsNoCacheCtx(ctx, &resp, query, userID)
		if err != nil {
			return nil, err
		}
		if len(resp) == 0 {
			err := redis.Setex(chaosSystemUserRoleUserIDKey, "*", int(unstable.AroundDuration(cacheOption.NotFoundExpiry).Seconds()))
			if err != nil {
				logx.Errorf("设置缓存失败, key: %v, error: %v", chaosSystemUserRoleUserIDKey, err)
			}
			return nil, ErrNotFound
		}
		if err := m.SetCacheCtx(ctx, chaosSystemUserRoleUserIDKey, resp); err != nil {
			logx.Errorf("设置缓存失败, key: %v, error: %v", chaosSystemUserRoleUserIDKey, err)
		}

		return resp, nil
	}

	return nil, err
}
