package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAPIAndCasbinLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteAPIAndCasbinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAPIAndCasbinLogic {
	return &DeleteAPIAndCasbinLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteAPIAndCasbinLogic) DeleteAPIAndCasbin(in *pb.DeleteAPIAndCasbinRequest) (*pb.Empty, error) {
	_, err := l.svcCtx.SyncedEnforcer.RemoveFilteredPolicy(1, in.Api, in.Method)
	if err != nil {
		return nil, errorx.Wrap(err, "删除casbin策略失败")
	}

	err = l.svcCtx.APIModel.Delete(l.ctx, in.ID)
	if err != nil {
		return nil, errors.Wrap(err, "数据库错误")
	}

	return &pb.Empty{}, nil
}
