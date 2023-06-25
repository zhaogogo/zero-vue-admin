package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AlertRuleLabelsModel = (*customAlertRuleLabelsModel)(nil)
var (
	cacheChaosMonitoringAlertRuleLabelsAlertRuleIdPrefix = "cache:chaosMonitoring:alertRuleLabels:alert_rule_id:"
)

type (
	// AlertRuleLabelsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAlertRuleLabelsModel.
	AlertRuleLabelsModel interface {
		Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error
		FindByAlertRuleID(ctx context.Context, alertruleid uint64) ([]AlertRuleLabels, error)
		TransFindByAlertRuleID(ctx context.Context, session sqlx.Session, alertruleid uint64) ([]AlertRuleLabels, error)
	}

	customAlertRuleLabelsModel struct {
		*defaultAlertRuleLabelsModel
	}
)

// NewAlertRuleLabelsModel returns a model for the database table.
func NewAlertRuleLabelsModel(conn sqlx.SqlConn, c cache.CacheConf) AlertRuleLabelsModel {
	return &customAlertRuleLabelsModel{
		defaultAlertRuleLabelsModel: newAlertRuleLabelsModel(conn, c),
	}
}

func (m *defaultAlertRuleLabelsModel) FindByAlertRuleID(ctx context.Context, alertruleid uint64) ([]AlertRuleLabels, error) {
	chaosMonitoringAlertRuleLabelsAlertRuleIdKey := fmt.Sprintf("%s%v", cacheChaosMonitoringAlertRuleLabelsAlertRuleIdPrefix, alertruleid)
	var resp []AlertRuleLabels
	err := m.QueryRowCtx(ctx, &resp, chaosMonitoringAlertRuleLabelsAlertRuleIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `alert_rule_id` = ?", alertRuleLabelsRows, m.table)
		return conn.QueryRowsCtx(ctx, v, query, alertruleid)
	})
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (m *defaultAlertRuleLabelsModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *defaultAlertRuleLabelsModel) TransFindByAlertRuleID(ctx context.Context, session sqlx.Session, alertruleid uint64) ([]AlertRuleLabels, error) {
	chaosMonitoringAlertRuleLabelsAlertRuleIdKey := fmt.Sprintf("%s%v", cacheChaosMonitoringAlertRuleLabelsAlertRuleIdPrefix, alertruleid)
	var resp []AlertRuleLabels
	err := m.QueryRowCtx(ctx, &resp, chaosMonitoringAlertRuleLabelsAlertRuleIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `alert_rule_id` = ?", alertRuleLabelsRows, m.table)
		return session.QueryRowsCtx(ctx, v, query, alertruleid)
	})
	if err != nil {
		return nil, err
	}

	return resp, err
}
