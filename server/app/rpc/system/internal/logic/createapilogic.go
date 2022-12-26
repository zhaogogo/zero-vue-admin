package logic

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/model/system"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateAPILogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateAPILogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAPILogic {
	return &CreateAPILogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateAPILogic) CreateAPI(in *pb.CreateAPIRequest) (*pb.Empty, error) {

	_, err := l.svcCtx.APIModel.Insert(l.ctx, &system.Api{
		Id:       0,
		Api:      in.API,
		Group:    in.Group,
		Describe: in.Describe,
		Method:   in.Method,
	})
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
