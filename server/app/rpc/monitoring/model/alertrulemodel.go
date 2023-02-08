package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ AlertRuleModel = (*customAlertRuleModel)(nil)

type (
	// AlertRuleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAlertRuleModel.
	AlertRuleModel interface {
		alertRuleModel
		FindPading_NC(ctx context.Context, param PagingAlertRuleParam) ([]AlertRule, error)
	}

	customAlertRuleModel struct {
		*defaultAlertRuleModel
	}
)

// NewAlertRuleModel returns a model for the database table.
func NewAlertRuleModel(conn sqlx.SqlConn, c cache.CacheConf) AlertRuleModel {
	return &customAlertRuleModel{
		defaultAlertRuleModel: newAlertRuleModel(conn, c),
	}
}

type PagingAlertRuleParam struct {
	Page     int64
	PageSize int64
	OrderKey string
	Order    string

	Search struct {
		Name string
		Type string
	}
}

func (m *defaultAlertRuleModel) FindPading_NC(ctx context.Context, param PagingAlertRuleParam) ([]AlertRule, error) {
	var resp []AlertRule
	query := ""
	orderSqlText := ""
	filtersql := []string{}
	if param.Search.Name != "" {
		filtersql = append(filtersql, fmt.Sprintf("`name` LIKE \"%%%s%%\"", param.Search.Name))
	}
	if param.Search.Type != "" {
		filtersql = append(filtersql, fmt.Sprintf("`type` LIKE \"%%%s%%\"", param.Search.Type))
	}
	if param.OrderKey != "" && param.Order == "ascending" {
		orderSqlText = fmt.Sprintf("ORDER BY `%s` asc", param.OrderKey)
	}
	if param.OrderKey != "" && param.Order == "descending" {
		orderSqlText = fmt.Sprintf("ORDER BY `%s` desc", param.OrderKey)
	}
	if len(filtersql) != 0 {
		query = fmt.Sprintf("select %s from %s WHERE %s %s LIMIT %d OFFSET %d", alertRuleRows, m.table, strings.Join(filtersql, " AND "), orderSqlText, param.PageSize, (param.Page-1)*param.PageSize)
	} else {
		query = fmt.Sprintf("select %s from %s %s LIMIT %d OFFSET %d", alertRuleRows, m.table, orderSqlText, param.PageSize, (param.Page-1)*param.PageSize)
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
