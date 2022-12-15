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
	cacheChaosSystemUserMenuParamsUserIdPrefix = "cache:chaosSystem:userMenuParams:user_id:"
)
var _ UserMenuParamsModel = (*customUserMenuParamsModel)(nil)

type (
	// UserMenuParamsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserMenuParamsModel.
	UserMenuParamsModel interface {
		userMenuParamsModel
		FindByUserID(ctx context.Context, redis *redis.Redis, userID uint64) ([]UserMenuParams, error)
	}

	customUserMenuParamsModel struct {
		*defaultUserMenuParamsModel
	}
)

// NewUserMenuParamsModel returns a model for the database table.
func NewUserMenuParamsModel(conn sqlx.SqlConn, c cache.CacheConf) UserMenuParamsModel {
	return &customUserMenuParamsModel{
		defaultUserMenuParamsModel: newUserMenuParamsModel(conn, c),
	}
}

func (m *defaultUserMenuParamsModel) FindByUserID(ctx context.Context, redis *redis.Redis, userID uint64) ([]UserMenuParams, error) {
	chaosSystemUserMenuParamsUserIdKey := fmt.Sprintf("%s%v", cacheChaosSystemUserMenuParamsUserIdPrefix, userID)
	var resp []UserMenuParams
	err := m.GetCacheCtx(ctx, chaosSystemUserMenuParamsUserIdKey, &resp)
	if err == nil {
		return resp, nil
	}
	if err.Error() == "placeholder" {
		return nil, ErrNotFound
	}
	if err == sql.ErrNoRows {
		query := fmt.Sprintf("select %s from %s where `user_id` = ?", userMenuParamsRows, m.table)
		err := m.QueryRowsNoCacheCtx(ctx, &resp, query, userID)
		if err != nil {
			return nil, err
		}
		if len(resp) == 0 {
			err := redis.Setex(chaosSystemUserMenuParamsUserIdKey, "*", int(unstable.AroundDuration(cacheOption.NotFoundExpiry).Seconds()))
			if err != nil {
				logx.Errorf("设置缓存失败, key: %v, error: %v", chaosSystemUserMenuParamsUserIdKey, err)
			}
			return nil, ErrNotFound
		}
		err = m.SetCacheCtx(ctx, chaosSystemUserMenuParamsUserIdKey, resp)
		if err != nil {
			logx.Errorf("设置缓存失败, key: %v, error: %v", chaosSystemUserMenuParamsUserIdKey, err)
		}
		return resp, nil
	}

	return nil, err
}
