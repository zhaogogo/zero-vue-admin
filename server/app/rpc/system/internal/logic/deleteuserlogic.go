package logic

import (
	"context"
	"database/sql"
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
		//删除用户
		err := l.svcCtx.UserModel.TransDelete(ctx, session, in.ID)
		if err != nil && err != sql.ErrNoRows {
			return err
		}

		//删除用户角色表
		err = l.svcCtx.UserRoleModel.TransDeleteByUserID(ctx, session, in.ID)
		if err != nil && err != sql.ErrNoRows {
			return err
		}

		//删除用户页面设置
		err = l.svcCtx.UserPageSetModel.TransDeleteByUserID(ctx, session, in.ID)
		if err != nil && err != sql.ErrNoRows {
			return err
		}

		//删除用户页面参数配置
		err = l.svcCtx.UserMenuParamsModel.TransDeleteByUserID(ctx, session, in.ID)
		if err != nil && err != sql.ErrNoRows {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
