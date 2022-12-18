package system

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MenuModel = (*customMenuModel)(nil)

type (
	// MenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMenuModel.
	MenuModel interface {
		menuModel
		FindAll_NC(ctx context.Context) ([]Menu, error)
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
