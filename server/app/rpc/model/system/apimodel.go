package system

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strconv"
	"strings"
)

var _ ApiModel = (*customApiModel)(nil)

type (
	// ApiModel is an interface to be customized, add more methods here,
	// and implement the added methods in customApiModel.
	ApiModel interface {
		apiModel
		Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error
		Total_NC(ctx context.Context, r *PagingAPIList) (int64, error)
		FindAll_NC(ctx context.Context) ([]Api, error)
		FindPaging_NC(ctx context.Context, r *PagingAPIList) ([]Api, error)

		TransUpdate(ctx context.Context, session sqlx.Session, newData *Api) error
		TransDeleteMultiple(ctx context.Context, session sqlx.Session, deleteMul []APIDeleteMultiple) error
	}

	customApiModel struct {
		*defaultApiModel
	}
)

// NewApiModel returns a model for the database table.
func NewApiModel(conn sqlx.SqlConn, c cache.CacheConf) ApiModel {
	return &customApiModel{
		defaultApiModel: newApiModel(conn, c),
	}
}

func (m *defaultApiModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *defaultApiModel) TransUpdate(ctx context.Context, session sqlx.Session, newData *Api) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	chaosSystemApiApiMethodKey := fmt.Sprintf("%s%v:%v", cacheChaosSystemApiApiMethodPrefix, data.Api, data.Method)
	chaosSystemApiIdKey := fmt.Sprintf("%s%v", cacheChaosSystemApiIdPrefix, data.Id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, apiRowsWithPlaceHolder)
		return session.ExecCtx(ctx, query, newData.Api, newData.Group, newData.Describe, newData.Method, newData.Id)
	}, chaosSystemApiApiMethodKey, chaosSystemApiIdKey)
	return err
}

type APIDeleteMultiple struct {
	ID     uint64
	API    string
	Method string
}

func (m *defaultApiModel) TransDeleteMultiple(ctx context.Context, session sqlx.Session, deleteMul []APIDeleteMultiple) error {
	cacheAllKey := []string{}
	ids := []string{}
	for _, d := range deleteMul {
		cacheAllKey = append(
			cacheAllKey,
			fmt.Sprintf("%s%v:%v", cacheChaosSystemApiApiMethodPrefix, d.API, d.Method),
			fmt.Sprintf("%s%v", cacheChaosSystemApiIdPrefix, d.ID),
		)
		ids = append(ids, strconv.FormatUint(d.ID, 10))
	}

	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` IN (%s)", m.table, strings.Join(ids, ","))
		return session.ExecCtx(ctx, query)
	}, cacheAllKey...)
	return err
}

func (m *defaultApiModel) FindAll_NC(ctx context.Context) ([]Api, error) {
	var resp []Api
	query := fmt.Sprintf("SELECT %s FROM %s", apiRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return nil, ErrNotFound
	}
	return resp, err
}

type PagingAPIList struct {
	Page     int64
	PageSize int64
	OrderKey string
	Order    string

	Api      string
	Describe string
	Group    string
	Method   string
}

func (m *defaultApiModel) FindPaging_NC(ctx context.Context, r *PagingAPIList) ([]Api, error) {
	var resp []Api
	ordersql := ""
	query := ""
	filtersql := []string{}
	if r.Api != "" {
		filtersql = append(filtersql, fmt.Sprintf("`api` LIKE \"%%%s%%\"", r.Api))
	}
	if r.Describe != "" {
		filtersql = append(filtersql, fmt.Sprintf("`describe` LIKE \"%%%s%%\"", r.Describe))
	}
	if r.Group != "" {
		filtersql = append(filtersql, fmt.Sprintf("`group` LIKE \"%%%s%%\"", r.Group))
	}
	if r.Method != "" {
		filtersql = append(filtersql, fmt.Sprintf("`method` = \"%s\"", r.Method))
	}
	if r.OrderKey != "" && r.Order == "ascending" {
		ordersql = fmt.Sprintf("ORDER BY `%s` asc", r.OrderKey)
	}
	if r.OrderKey != "" && r.Order == "descending" {
		ordersql = fmt.Sprintf("ORDER BY `%s` desc", r.OrderKey)
	}
	if len(filtersql) != 0 {
		query = fmt.Sprintf("SELECT %s FROM %s WHERE %s %s LIMIT %d OFFSET %d", apiRows, m.table, strings.Join(filtersql, " AND "), ordersql, r.PageSize, (r.Page-1)*r.PageSize)
	} else {
		query = fmt.Sprintf("SELECT %s FROM %s %s LIMIT %d OFFSET %d", apiRows, m.table, ordersql, r.PageSize, (r.Page-1)*r.PageSize)
	}
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return nil, ErrNotFound
	}
	return resp, err
}

func (m *defaultApiModel) Total_NC(ctx context.Context, r *PagingAPIList) (int64, error) {
	var resp int64
	query := ""
	filtersql := []string{}
	if r.Api != "" {
		filtersql = append(filtersql, fmt.Sprintf("`api` LIKE \"%%%s%%\"", r.Api))
	}
	if r.Describe != "" {
		filtersql = append(filtersql, fmt.Sprintf("`describe` LIKE \"%%%s%%\"", r.Describe))
	}
	if r.Group != "" {
		filtersql = append(filtersql, fmt.Sprintf("`group` LIKE \"%%%s%%\"", r.Group))
	}
	if r.Method != "" {
		filtersql = append(filtersql, fmt.Sprintf("`method` = \"%s\"", r.Method))
	}

	if len(filtersql) != 0 {
		query = fmt.Sprintf("SELECT count(*) AS total FROM %s WHERE %s", m.table, strings.Join(filtersql, " AND "))
	} else {
		query = fmt.Sprintf("SELECT count(*) AS total FROM %s", m.table)
	}

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
