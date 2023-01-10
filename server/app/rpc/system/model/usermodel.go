package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
	"strings"
)

var _ UserModel = (*customUserModel)(nil)
var userRowsWithPlaceHolderWithOutPassword = strings.Join(stringx.Remove(userFieldNames, "`id`", "`password`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		TransCtx(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error
		FindPagingList_NC(ctx context.Context, r *PagingUserList) ([]User, error)
		FindAll_NC(ctx context.Context) ([]User, error)
		Total_NC(ctx context.Context) (int64, error)

		FindOneByNameWHEREDeleteTimeISNULL(ctx context.Context, name string) (*User, error)

		UpdateUserPassword(ctx context.Context, id uint64, pass string) error
		UpdateDeleteColumn(ctx context.Context, userid uint64, deleteby string, deletetime sql.NullTime) error
		UpdateWithOutPassword(ctx context.Context, newData *User) error
		UpdateCurrentRoleColumn(ctx context.Context, userid uint64, roleid uint64) error

		TransInsert(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error)
		TransDelete(ctx context.Context, session sqlx.Session, id uint64) error
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

func (m *defaultUserModel) TransDelete(ctx context.Context, session sqlx.Session, id uint64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	chaosSystemUserIdKey := fmt.Sprintf("%s%v", cacheChaosSystemUserIdPrefix, id)
	chaosSystemUserNameKey := fmt.Sprintf("%s%v", cacheChaosSystemUserNamePrefix, data.Name)
	chaossystemUserMode_userId_key := fmt.Sprintf("%v%v", cacheChaosSystemUserRoleUserIdPrefix, id) //删除user_role中间表
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return session.ExecCtx(ctx, query, id)
	}, chaosSystemUserIdKey, chaosSystemUserNameKey, chaossystemUserMode_userId_key)
	return err
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

type PagingUserList struct {
	Page     int64
	PageSize int64
	NameX    string
}

func (m *defaultUserModel) FindPagingList_NC(ctx context.Context, r *PagingUserList) ([]User, error) {
	var (
		resp  []User
		query = ""
	)
	if r.NameX == "" {
		query = fmt.Sprintf("SELECT %s FROM %s LIMIT %d OFFSET %d", userRows, m.table, r.PageSize, (r.Page-1)*r.PageSize)
	} else {
		query = fmt.Sprintf("SELECT %s FROM %s WHERE `name` LIKE \"%%%s%%\" OR `nick_name` LIKE \"%%%s%%\" OR `email` LIKE \"%%%s%%\" LIMIT %d OFFSET %d", userRows, m.table, r.NameX, r.NameX, r.NameX, r.PageSize, (r.Page-1)*r.PageSize)
	}
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return nil, ErrNotFound
	}

	return resp, nil
}

func (m *defaultUserModel) FindAll_NC(ctx context.Context) ([]User, error) {
	var resp []User
	query := fmt.Sprintf("SELECT %s FROM %s", userRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return nil, ErrNotFound
	}

	return resp, nil
}

func (m *defaultUserModel) Total_NC(ctx context.Context) (int64, error) {
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

func (m *defaultUserModel) UpdateWithOutPassword(ctx context.Context, newData *User) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	chaosSystemUserIdKey := fmt.Sprintf("%s%v", cacheChaosSystemUserIdPrefix, data.Id)
	chaosSystemUserNameKey := fmt.Sprintf("%s%v", cacheChaosSystemUserNamePrefix, data.Name)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolderWithOutPassword)
		return conn.ExecCtx(ctx, query, newData.Name, newData.NickName, data.Type, newData.Email, newData.Phone, newData.Department, newData.Position, data.CreateBy, newData.UpdateBy, data.DeleteBy, data.DeleteTime, data.PageSetId, newData.Id)
	}, chaosSystemUserIdKey, chaosSystemUserNameKey)
	return err
}

func (m *defaultUserModel) UpdateCurrentRoleColumn(ctx context.Context, userid uint64, roleid uint64) error {
	chaosSystemUserIdKey := fmt.Sprintf("%s%v", cacheChaosSystemUserIdPrefix, userid)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `current_role` = ? where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, roleid, userid)
	}, chaosSystemUserIdKey)
	if err != nil {
		return err
	}
	return nil
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
	//软删除 需要删除中间表缓存
	cacheChaosSystemUserRole_UserId_Key := fmt.Sprintf("%v%v", cacheChaosSystemUserRoleUserIdPrefix, userid)

	chaosSystemUserIdKey := fmt.Sprintf("%s%v", cacheChaosSystemUserIdPrefix, userid)
	chaossystemUserMode_userId_key := fmt.Sprintf("%v%v", cacheChaosSystemUserRoleUserIdPrefix, userid) //删除user_role中间表
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("UPDATE %s set `delete_by` = ?, `delete_time` = ? where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, deleteby, deletetime, userid)
	}, chaosSystemUserIdKey, chaossystemUserMode_userId_key, cacheChaosSystemUserRole_UserId_Key)
	return err
}
