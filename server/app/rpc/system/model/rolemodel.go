package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RoleModel = (*customRoleModel)(nil)

type (
	// RoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRoleModel.
	RoleModel interface {
		roleModel
		FindAll_NC(ctx context.Context) ([]Role, error)
		Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error

		TransDelete(ctx context.Context, session sqlx.Session, id uint64) error
		UpdateDeleteColumn(ctx context.Context, roleid uint64, deleteby string, deletetime sql.NullTime) error
	}

	customRoleModel struct {
		*defaultRoleModel
	}
)

// NewRoleModel returns a model for the database table.
func NewRoleModel(conn sqlx.SqlConn, c cache.CacheConf) RoleModel {
	return &customRoleModel{
		defaultRoleModel: newRoleModel(conn, c),
	}
}

func (m *defaultRoleModel) FindAll_NC(ctx context.Context) ([]Role, error) {
	var resp []Role
	query := fmt.Sprintf("SELECT %s FROM %s", roleRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return nil, ErrNotFound
	}
	return resp, nil
}
func (m *defaultRoleModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *defaultRoleModel) TransDelete(ctx context.Context, session sqlx.Session, id uint64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	chaosSystemRoleIdKey := fmt.Sprintf("%s%v", cacheChaosSystemRoleIdPrefix, id)
	chaosSystemRoleRoleKey := fmt.Sprintf("%s%v", cacheChaosSystemRoleRolePrefix, data.Role)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return session.ExecCtx(ctx, query, id)
	}, chaosSystemRoleIdKey, chaosSystemRoleRoleKey)
	return err
}

func (m *defaultRoleModel) UpdateDeleteColumn(ctx context.Context, roleid uint64, deleteby string, deletetime sql.NullTime) error {
	//软删除 需要删除中间表缓存
	allKey := []string{}
	// 删除role_menu缓存
	allKey = append(allKey, fmt.Sprintf("%s%v", cacheChaosSystemRoleMenuRoleIdPrefix, roleid)) // 删除role_menu缓存
	// 删除user_role缓存
	userroles := []UserRole{}
	useridset := make(map[uint64]int64)
	q := fmt.Sprintf("SELECT `id`,`user_id`,`role_id` FROM `user_role` WHERE `role_id` = ?")
	err := m.QueryRowsNoCacheCtx(ctx, &userroles, q, roleid)
	if err != nil {
		return err
	}
	for _, userrole := range userroles {
		useridset[userrole.UserId] += 1
	}

	for userid, _ := range useridset {
		allKey = append(allKey, fmt.Sprintf("%v%v", cacheChaosSystemUserRoleUserIdPrefix, userid))
	}
	//删除本地缓存
	allKey = append(allKey, fmt.Sprintf("%s%v", cacheChaosSystemRoleIdPrefix, roleid))
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("UPDATE %s set `delete_by` = ?, `delete_time` = ? where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, deleteby, deletetime, roleid)
	}, allKey...)
	return err
}
