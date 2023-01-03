package logic

import (
	"context"
	"github.com/pkg/errors"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCasbinPolicyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCasbinPolicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCasbinPolicyLogic {
	return &UpdateCasbinPolicyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCasbinPolicyLogic) UpdateCasbinPolicy(in *pb.UpdateCasbinPolicyRequest) (*pb.Empty, error) {
	l.svcCtx.SyncedEnforcer.RemoveFilteredPolicy(0, in.V0)
	rules := [][]string{}
	for _, v := range in.CasbinRules {
		rules = append(rules, []string{in.V0, v.V1, v.V2})
	}
	success, err := l.svcCtx.SyncedEnforcer.AddPolicies(rules)
	if err != nil {
		return nil, errors.Wrap(err, "添加piolicy失败")
	}
	if !success {
		return nil, errors.New("存在相同API policy,联系管理员")
	}
	return &pb.Empty{}, nil
}
