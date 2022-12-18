package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserRoleLogic {
	return &UpdateUserRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserRoleLogic) UpdateUserRole(in *pb.UpdateUserRoleRequest) (*pb.Empty, error) {
	err := l.svcCtx.UserRoleModel.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		err := l.svcCtx.UserRoleModel.TransDeleteByUserID(ctx, session, in.UserID)
		if err != nil {
			return errors.Wrap(err, "删除用户权限失败")
		}

		if len(in.RoleList) == 0 {
			return nil
		}

		rest, err := l.svcCtx.UserRoleModel.TranInsertUserIDRoleIDs(l.ctx, session, in.UserID, in.RoleList)
		if err != nil {
			return errors.Wrap(err, "新增用户权限失败")
		}
		a, err := rest.RowsAffected()
		if err != nil {
			return errors.Wrap(err, "获取插入行数错误")
		}
		if a != int64(len(in.RoleList)) {
			return errors.New("插入行数不匹配")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
