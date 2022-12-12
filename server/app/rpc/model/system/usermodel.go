package system

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		FindOneByNameWHEREDeleteTimeISNULL(ctx context.Context, name string) (*User, error)
		FindListPaging(ctx context.Context, page int64, pageSize int64) ([]User, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c),
	}
}

func (m *defaultUserModel) FindOneByNameWHEREDeleteTimeISNULL(ctx context.Context, name string) (*User, error) {
	cacheChaosSystemUserNameKey := fmt.Sprintf("%s%v", cacheChaosSystemUserNamePrefix, name)
	var resp User
	err := m.QueryRowIndexCtx(ctx, &resp, cacheChaosSystemUserNameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `name` = ? AND `delete_time` IS NULL limit 1", userRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, name); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindListPaging(ctx context.Context, page int64, pageSize int64) ([]User, error) {
	var resp []User
	query := fmt.Sprintf("SELECT %s FROM %s LIMIT %d OFFSET %d", userRows, m.table, pageSize, (page-1)*pageSize)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return nil, ErrNotFound
	}

	return resp, nil
}
