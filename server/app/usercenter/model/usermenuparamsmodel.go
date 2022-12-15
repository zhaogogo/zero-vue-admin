package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var (
	cacheUsercenterUserMenuParamsUserIdPrefix = "cache:usercenter:userMenuParams:user_id:"
)

var _ UserMenuParamsModel = (*customUserMenuParamsModel)(nil)

type (
	// UserMenuParamsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserMenuParamsModel.
	UserMenuParamsModel interface {
		userMenuParamsModel
		FindByUserID(ctx context.Context, userid uint64) ([]UserMenuParams, error)
		FindAll(ctx context.Context) ([]UserMenuParams, error)
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

func (m *defaultUserMenuParamsModel) FindByUserID(ctx context.Context, userid uint64) ([]UserMenuParams, error) {
	usercenterUserMenuParamsUserIDKey := fmt.Sprintf("%s%v", cacheUsercenterUserMenuParamsUserIdPrefix, userid)
	var resp []UserMenuParams
	err := m.QueryRowCtx(ctx, &resp, usercenterUserMenuParamsUserIDKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("SELECT %s FROM %s WHERE `user_id` = ?", userMenuParamsRows, m.table)
		return conn.QueryRowsCtx(ctx, v, query, userid)
	})
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserMenuParamsModel) FindAll(ctx context.Context) ([]UserMenuParams, error) {
	var resp []UserMenuParams
	query := fmt.Sprintf("SELECT %s FROM %s", userMenuParamsRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	fmt.Println("userMenuParams", err)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
