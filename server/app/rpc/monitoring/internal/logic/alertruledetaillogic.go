package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/monitoring/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/monitoring/model"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/monitoring/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlertRuleDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAlertRuleDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlertRuleDetailLogic {
	return &AlertRuleDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AlertRuleDetailLogic) AlertRuleDetail(in *pb.AlertRuleID) (*pb.AlertRuleDetailResponse, error) {
	res := &pb.AlertRuleDetailResponse{Labels: make(map[string]string)}
	l.svcCtx.AlertRuleModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		alertrule, err := l.svcCtx.AlertRuleModel.FindOne(l.ctx, in.ID)
		if err != nil {
			if err == model.ErrNotFound {
				return err
			}
			return errors.Wrap(err, "数据库[alert_rule]查询错误")
		}
		res.ID = alertrule.Id
		res.Name = alertrule.Name
		res.Type = alertrule.Type
		res.Group = alertrule.Group
		res.Tag = alertrule.Tag
		res.To = alertrule.To
		res.Expr = alertrule.Expr
		res.Operator = alertrule.Operator
		res.Value = alertrule.Value
		res.For = alertrule.For
		res.Summary = alertrule.AnnoSummary
		res.Describe = alertrule.AnnoDesc
		res.IsWrite = alertrule.IsWrite == 1

		labels, err := l.svcCtx.AlertRuleLabels.TransFindByAlertRuleID(ctx, session, res.ID)
		if err != nil {
			return errors.Wrap(err, "数据库[alert_rule_labels]查询错误")
		}
		if labels != nil {
			for _, v := range labels {
				res.Labels[v.Key] = v.Value
			}
		}

		return nil
	})

	return res, nil
}
