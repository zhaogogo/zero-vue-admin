package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/monitoring/model"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/monitoring/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/monitoring/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlertRulePagingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAlertRulePagingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlertRulePagingLogic {
	return &AlertRulePagingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AlertRulePagingLogic) AlertRulePaging(in *pb.AlertRulePagingRequest) (*pb.AlertRulePagingResponse, error) {
	param := model.PagingAlertRuleParam{
		Page:     in.Page,
		PageSize: in.PageSize,
		OrderKey: in.OrderKey,
		Order:    in.Order,
		Search: struct {
			Name string
			Type string
		}{Name: in.Name, Type: in.Type},
	}
	AlertRules, err := l.svcCtx.AlertRuleModel.FindPading_NC(l.ctx, param)
	if err != nil {
		if err != model.ErrNotFound {
			return nil, err
		}
		return nil, errors.Wrapf(err, "数据库查询错误")
	}
	pAlertRules := []*pb.AlertRule{}
	for _, v := range AlertRules {
		pAlertRule := &pb.AlertRule{
			ID:       v.Id,
			Name:     v.Name,
			Type:     v.Type,
			Group:    v.Group,
			Tag:      v.Tag,
			To:       v.To,
			Expr:     v.Expr,
			Operator: v.Operator,
			Value:    v.Value,
			For:      v.For,
			IsWrite:  v.IsWrite == 1,
		}
		if v.Value == "" {
			pAlertRule.Describe = v.Describe
			pAlertRule.Summary = v.Summary
		} else {
			pAlertRule.Describe = v.Describe + v.Operator + v.Value
			pAlertRule.Summary = v.Summary + v.Operator + v.Value
		}
		pAlertRules = append(pAlertRules, pAlertRule)
	}
	return &pb.AlertRulePagingResponse{
		AlertRules: pAlertRules,
	}, nil
}
