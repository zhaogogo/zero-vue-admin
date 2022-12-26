package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CasbinEnforcerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCasbinEnforcerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CasbinEnforcerLogic {
	return &CasbinEnforcerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CasbinEnforcerLogic) CasbinEnforcer(in *pb.CasbinEnforceRequest) (*pb.CasbinEnforceResponse, error) {
	success, err := l.svcCtx.SyncedEnforcer.Enforce(in.Sub, in.Obj, in.Act)
	if err != nil {
		return nil, errorx.Wrap(err, "权限认证错误")
	}
	return &pb.CasbinEnforceResponse{
		Success: success,
	}, nil
}
