package logic

import (
	"context"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/esManager/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/esManager/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteESConnLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteESConnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteESConnLogic {
	return &DeleteESConnLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteESConnLogic) DeleteESConn(in *pb.ESConnID) (*pb.Empty, error) {
	err := l.svcCtx.ESConnModel.Delete(l.ctx, in.ID)
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
