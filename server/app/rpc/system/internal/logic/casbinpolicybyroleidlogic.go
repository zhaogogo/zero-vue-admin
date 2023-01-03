package logic

import (
	"context"
	"strconv"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CasbinPolicyByRoleIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCasbinPolicyByRoleIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CasbinPolicyByRoleIDLogic {
	return &CasbinPolicyByRoleIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CasbinPolicyByRoleIDLogic) CasbinPolicyByRoleID(in *pb.RoleID) (*pb.CasbinPolicyResponse, error) {
	pbPolicy := []*pb.CasbinPolicy{}
	policy := l.svcCtx.SyncedEnforcer.GetFilteredPolicy(0, strconv.FormatUint(in.ID, 10))
	for _, p := range policy {
		pbPolicy = append(pbPolicy, &pb.CasbinPolicy{Api: p[1], Method: p[2]})
	}

	return &pb.CasbinPolicyResponse{
		Policy: pbPolicy,
	}, nil
}
