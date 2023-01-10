package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MenuModel = (*customMenuModel)(nil)

type (
	// MenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMenuModel.
	MenuModel interface {
		menuModel
		Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error
		FindAll_NC(ctx context.Context) ([]Menu, error)
		FindByParentID_NC(ctx context.Context, parentid uint64) ([]Menu, error)

		TransDelete(ctx context.Context, session sqlx.Session, id uint64) error
	}

	customMenuModel struct {
		*defaultMenuModel
	}
)

// NewMenuModel returns a model for the database table.
func NewMenuModel(conn sqlx.SqlConn, c cache.CacheConf) MenuModel {
	return &customMenuModel{
		defaultMenuModel: newMenuModel(conn, c),
	}
}
func (m *defaultMenuModel) FindByParentID_NC(ctx context.Context, parentid uint64) ([]Menu, error) {
	var resp []Menu
	query := fmt.Sprintf("select %s from %s where `parent_id` = ?", menuRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, parentid)
	if err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return nil, ErrNotFound
	}

	return resp, nil
}
func (m *defaultMenuModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *defaultMenuModel) TransDelete(ctx context.Context, session sqlx.Session, id uint64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	if _, err = m.FindByParentID_NC(ctx, data.Id); err != nil && err != ErrNotFound {
		return err //数据库错误
	} else if err == ErrNotFound {
		//真实删除每次都会删除中间表
		chaosSystemMenuIdKey := fmt.Sprintf("%s%v", cacheChaosSystemMenuIdPrefix, id)
		chaosSystemMenuNameKey := fmt.Sprintf("%s%v", cacheChaosSystemMenuNamePrefix, data.Name)
		chaosSystemMenuPathKey := fmt.Sprintf("%s%v", cacheChaosSystemMenuPathPrefix, data.Path)
		_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
			query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
			return session.ExecCtx(ctx, query, id)
		}, chaosSystemMenuIdKey, chaosSystemMenuNameKey, chaosSystemMenuPathKey)
		return err
	} else {
		return errors.New("存在子菜单不可删除")
	}

}

func (m *defaultMenuModel) FindAll_NC(ctx context.Context) ([]Menu, error) {
	var resp []Menu
	query := fmt.Sprintf("select %s from %s", menuRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return nil, ErrNotFound
	}
	return resp, nil
}
