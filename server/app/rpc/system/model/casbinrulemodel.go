package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strconv"
	"strings"
)

var _ CasbinRuleModel = (*customCasbinRuleModel)(nil)

type (
	// CasbinRuleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCasbinRuleModel.
	CasbinRuleModel interface {
		casbinRuleModel
		TransUpdateV2V3(ctx context.Context, session sqlx.Session, newdata *Api, oldData *Api) error
		TransDeleteMultiple(ctx context.Context, session sqlx.Session, deleteMul []APIDeleteMultiple) error
		TransDeleteByV0(ctx context.Context, session sqlx.Session, roleid uint64) error
	}

	customCasbinRuleModel struct {
		*defaultCasbinRuleModel
	}
)

// NewCasbinRuleModel returns a model for the database table.
func NewCasbinRuleModel(conn sqlx.SqlConn) CasbinRuleModel {
	return &customCasbinRuleModel{
		defaultCasbinRuleModel: newCasbinRuleModel(conn),
	}
}

func (m *defaultCasbinRuleModel) TransUpdateV2V3(ctx context.Context, session sqlx.Session, newdata *Api, oldData *Api) error {
	query := fmt.Sprintf("UPDATE %s SET `v1` = ?, `v2` = ? WHERE `v1` = ? AND `v2` = ?", m.table)
	_, err := session.ExecCtx(ctx, query, newdata.Api, newdata.Method, oldData.Api, oldData.Method)
	return err
}

func (m *defaultCasbinRuleModel) TransDeleteByV0(ctx context.Context, session sqlx.Session, roleid uint64) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE `v0` = ?", m.table)
	_, err := session.ExecCtx(ctx, query, strconv.FormatUint(roleid, 10))
	if err != nil {
		return err
	}
	return nil
}

func (m *defaultCasbinRuleModel) TransDeleteMultiple(ctx context.Context, session sqlx.Session, deleteMul []APIDeleteMultiple) error {
	condition := []string{}
	for _, d := range deleteMul {
		condition = append(condition, "(`v1` = \""+d.API+"\" AND `v2` = \""+d.Method+"\")")
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE %s", m.table, strings.Join(condition, " OR "))
	_, err := session.ExecCtx(ctx, query)
	return err
}
