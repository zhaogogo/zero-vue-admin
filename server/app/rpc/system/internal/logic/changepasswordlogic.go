package logic

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/common/utils"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordLogic {
	return &ChangePasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChangePasswordLogic) ChangePassword(in *pb.ChangePasswordRequest) (*pb.Empty, error) {
	pass, err := utils.GenPassword(in.Password)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.UserModel.UpdateUserPassword(l.ctx, in.ID, pass)
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
