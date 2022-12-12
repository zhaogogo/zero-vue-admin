package system

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
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
		FindByUserID(ctx context.Context, userID uint64) ([]UserMenuParams, error)
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

func (m *defaultUserMenuParamsModel) FindByUserID(ctx context.Context, userID uint64) ([]UserMenuParams, error) {
	chaosSystemUserMenuParamsUserIdKey := fmt.Sprintf("%s%v", cacheChaosSystemUserMenuParamsUserIdPrefix, userID)
	var resp []UserMenuParams
	query := fmt.Sprintf("select %s from %s where `user_id` = ?", userMenuParamsRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, userID)
	if err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return nil, ErrNotFound
	}
	m.SetCacheCtx(ctx, chaosSystemUserMenuParamsUserIdKey, resp)
	return resp, nil
}
