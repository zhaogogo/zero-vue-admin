package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserLogic) DeleteUser(in *pb.UserID) (*pb.Empty, error) {
	err := l.svcCtx.UserModel.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		err := l.svcCtx.UserModel.TransDelete(ctx, session, in.ID)
		if err != nil {
			return err
		}

		err = l.svcCtx.UserRoleModel.TransDeleteByUserID(ctx, session, in.ID)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
