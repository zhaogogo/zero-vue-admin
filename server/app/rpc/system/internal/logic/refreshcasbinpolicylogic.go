package logic

import (
	"context"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshCasbinPolicyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRefreshCasbinPolicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshCasbinPolicyLogic {
	return &RefreshCasbinPolicyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RefreshCasbinPolicyLogic) RefreshCasbinPolicy(in *pb.Empty) (*pb.Empty, error) {
	err := l.svcCtx.SyncedEnforcer.LoadPolicy()
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
