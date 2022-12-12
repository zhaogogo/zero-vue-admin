package model

import (
	"context"
	"database/sql"
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
		FindAll(ctx context.Context) ([]User, error)
		FindPaging(ctx context.Context, page int, pageSize int) ([]User, error)
		FindOneByNameWHEREDeleteTimeISNULL(ctx context.Context, name string) (*User, error)
		FindOneByIDWHEREDeleteTimeISNULL(ctx context.Context, id uint64) (*User, error)
		UpdateDeleteColumn(ctx context.Context, userid uint64, deleteby string, deletetime sql.NullTime) error
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

func (m *defaultUserModel) FindAll(ctx context.Context) ([]User, error) {
	var resp []User
	query := fmt.Sprintf("SELECT %s FROM %s", userRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindPaging(ctx context.Context, page int, pageSize int) ([]User, error) {
	var resp []User
	query := fmt.Sprintf("SELECT %s FROM %s LIMIT %v OFFSET %v", userRows, m.table, pageSize, (page-1)*pageSize)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByNameWHEREDeleteTimeISNULL(ctx context.Context, name string) (*User, error) {
	usercenterUserNameKey := fmt.Sprintf("%s%v", cacheUsercenterUserNamePrefix, name)
	var resp User
	err := m.QueryRowIndexCtx(ctx, &resp, usercenterUserNameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
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

func (m *defaultUserModel) FindOneByIDWHEREDeleteTimeISNULL(ctx context.Context, id uint64) (*User, error) {
	usercenterUserIdKey := fmt.Sprintf("%s%v", cacheUsercenterUserIdPrefix, id)
	var resp User
	err := m.QueryRowCtx(ctx, &resp, usercenterUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? AND `delete_time` IS NULL limit 1", userRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) UpdateDeleteColumn(ctx context.Context, userid uint64, deleteby string, deletetime sql.NullTime) error {
	data, err := m.FindOne(ctx, userid)
	if err != nil {
		return err
	}

	usercenterUserIdKey := fmt.Sprintf("%s%v", cacheUsercenterUserIdPrefix, data.Id)
	usercenterUserNameKey := fmt.Sprintf("%s%v", cacheUsercenterUserNamePrefix, data.Name)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `delete_by` = ?, `delete_time` = ? where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, deleteby, deletetime, userid)
	}, usercenterUserIdKey, usercenterUserNameKey)
	return err
}
