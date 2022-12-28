package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleMenusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleMenusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleMenusLogic {
	return &UpdateRoleMenusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateRoleMenusLogic) UpdateRoleMenus(in *pb.UpdateRoleMenusRequest) (*pb.Empty, error) {
	err := l.svcCtx.RoleMenuModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		return l.svcCtx.RoleMenuModel.TransReplaceByMenus(ctx, session, in.RoleID, in.MenuIDList)
	})
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
