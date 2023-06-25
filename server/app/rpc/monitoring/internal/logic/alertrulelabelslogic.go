package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/monitoring/model"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/monitoring/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/monitoring/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlertRuleLabelsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAlertRuleLabelsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlertRuleLabelsLogic {
	return &AlertRuleLabelsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AlertRuleLabelsLogic) AlertRuleLabels(in *pb.AlertRuleID) (*pb.AlertRuleLabelsResponse, error) {
	label := make(map[string]string, 3)
	alertrule, err := l.svcCtx.AlertRuleLabels.FindByAlertRuleID(l.ctx, in.ID)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, err
		}
		return nil, errors.Wrap(err, "数据库[alert_rule_labels]查询错误")
	}
	for _, v := range alertrule {
		label[v.Key] = v.Value
	}
	return &pb.AlertRuleLabelsResponse{Labels: label}, nil
}
