package logic

import (
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMenuRoleMenuUserMenuParamLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMenuRoleMenuUserMenuParamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMenuRoleMenuUserMenuParamLogic {
	return &DeleteMenuRoleMenuUserMenuParamLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteMenuRoleMenuUserMenuParamLogic) DeleteMenu_RoleMenu_UserMenuParam(in *pb.MenuID) (*pb.Empty, error) {
	err := l.svcCtx.MenuModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		//删除菜单
		err := l.svcCtx.MenuModel.TransDelete(ctx, session, in.ID)
		if err != nil && err != sql.ErrNoRows {
			return err
		}

		//删除角色菜单权限
		err = l.svcCtx.RoleMenuModel.TransDeleteByMenuId(ctx, session, in.ID)
		if err != nil {
			return err
		}

		//删除用户菜单参数
		err = l.svcCtx.UserMenuParamsModel.TransDeleteByMenuID(ctx, session, in.ID)
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
