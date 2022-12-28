package logic

import (
	"context"
	"github.com/pkg/errors"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserCurrentRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserCurrentRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserCurrentRoleLogic {
	return &UpdateUserCurrentRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserCurrentRoleLogic) UpdateUserCurrentRole(in *pb.UpdateUserCurrentRoleRequest) (*pb.Empty, error) {
	err := l.svcCtx.UserModel.UpdateCurrentRoleColumn(l.ctx, in.UserID, in.RoleID)
	if err != nil {
		return nil, errors.Wrap(err, "数据库错误")
	}
	return &pb.Empty{}, nil
}
