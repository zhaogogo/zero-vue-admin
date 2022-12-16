package system

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
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
		Total(ctx context.Context) (int64, error)
		UpdateUserPassword(ctx context.Context, id uint64, pass string) error
		UpdateDeleteColumn(ctx context.Context, userid uint64, deleteby string, deletetime sql.NullTime) error

		TransCtx(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error
		TransInsert(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error)
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

func (m *defaultUserModel) TransCtx(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *defaultUserModel) TransInsert(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error) {
	chaosSystemUserIdKey := fmt.Sprintf("%s%v", cacheChaosSystemUserIdPrefix, data.Id)
	chaosSystemUserNameKey := fmt.Sprintf("%s%v", cacheChaosSystemUserNamePrefix, data.Name)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
		return session.ExecCtx(ctx, query, data.Name, data.NickName, data.Password, data.Type, data.Email, data.Phone, data.Department, data.Position, data.CreateBy, data.UpdateBy, data.DeleteBy, data.DeleteTime, data.PageSetId)
	}, chaosSystemUserIdKey, chaosSystemUserNameKey)
	return ret, err
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

func (m *defaultUserModel) Total(ctx context.Context) (int64, error) {
	var total int64
	query := fmt.Sprintf("SELECT count(*) AS total FROM %s", m.table)
	err := m.QueryRowNoCache(&total, query)
	switch err {
	case nil:
		return total, nil
	case ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (m *defaultUserModel) UpdateUserPassword(ctx context.Context, id uint64, pass string) error {
	chaosSystemUserIdKey := fmt.Sprintf("%s%v", cacheChaosSystemUserIdPrefix, id)
	res, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `password` = ? where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, pass, id)
	}, chaosSystemUserIdKey)
	if err != nil {
		return err
	}

	rowsAffect, err := res.RowsAffected()
	if err != nil {
		logx.Error("修改密码获取RowsAffected失败")
	}
	if rowsAffect != 1 {
		logx.Error("修改密码影响数据超过1")
	}
	return nil
}

func (m *defaultUserModel) UpdateDeleteColumn(ctx context.Context, userid uint64, deleteby string, deletetime sql.NullTime) error {
	chaosSystemUserIdKey := fmt.Sprintf("%s%v", cacheChaosSystemUserIdPrefix, userid)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("UPDATE %s set `delete_by` = ?, `delete_time` = ? where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, deleteby, deletetime, userid)
	}, chaosSystemUserIdKey)
	return err
}
