package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoleLogic {
	return &DeleteRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteRoleLogic) DeleteRole(in *pb.RoleID) (*pb.Empty, error) {
	err := l.svcCtx.RoleModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		// 删除role模型
		err := l.svcCtx.RoleModel.TransDelete(ctx, session, in.ID)
		if err != nil {
			return errors.Wrap(err, "删除role失败")
		}
		// 删除user_role模型
		err = l.svcCtx.UserRoleModel.TransDeleteByRoleID(ctx, session, in.ID)
		if err != nil {
			return errors.Wrap(err, "删除user_role失败")
		}
		// 删除role_menu模型
		err = l.svcCtx.RoleMenuModel.TransDeleteByRoleId(ctx, session, in.ID)
		if err != nil {
			return errors.Wrap(err, "删除role_menu失败")
		}

		//删除casbin_rule规则
		err = l.svcCtx.CasbinRuleModel.TransDeleteByV0(ctx, session, in.ID)
		if err != nil {
			return errors.Wrap(err, "删除casbin_rule失败")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	l.svcCtx.SyncedEnforcer.LoadPolicy()
	return &pb.Empty{}, nil
}
