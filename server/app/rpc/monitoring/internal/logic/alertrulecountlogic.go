package logic

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/monitoring/model"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/monitoring/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/monitoring/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlertRuleCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAlertRuleCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlertRuleCountLogic {
	return &AlertRuleCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AlertRuleCountLogic) AlertRuleCount(in *pb.AlertRuleCountRequest) (*pb.Total, error) {
	param := model.CountAlertRuleParam{
		Search: struct {
			Name string
			Type string
		}{
			Name: in.Name,
			Type: in.Type,
		},
	}
	total, err := l.svcCtx.AlertRuleModel.Count_NC(l.ctx, param)
	if err != nil {
		return nil, err
	}

	return &pb.Total{
		Total: total,
	}, nil
}
