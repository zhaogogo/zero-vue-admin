package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ EsConnModel = (*customEsConnModel)(nil)

type (
	// EsConnModel is an interface to be customized, add more methods here,
	// and implement the added methods in customEsConnModel.
	EsConnModel interface {
		esConnModel
		FindPaging_NC(ctx context.Context, page int64, pageSize int64) ([]EsConn, error)
		FindTotal_NC(ctx context.Context) (int64, error)
	}

	customEsConnModel struct {
		*defaultEsConnModel
	}
)

// NewEsConnModel returns a model for the database table.
func NewEsConnModel(conn sqlx.SqlConn, c cache.CacheConf) EsConnModel {
	return &customEsConnModel{
		defaultEsConnModel: newEsConnModel(conn, c),
	}
}

func (m *defaultEsConnModel) FindPaging_NC(ctx context.Context, page int64, pageSize int64) ([]EsConn, error) {
	var resp []EsConn
	query := fmt.Sprintf("select %s from %s limit ? offset ?", esConnRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, pageSize, (page-1)*pageSize)
	if err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return nil, ErrNotFound
	}
	return resp, nil
}

func (m *defaultEsConnModel) FindTotal_NC(ctx context.Context) (int64, error) {
	var resp int64
	query := fmt.Sprintf("SELECT count(*) AS total FROM %s", m.table)
	err := m.QueryRowNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}
