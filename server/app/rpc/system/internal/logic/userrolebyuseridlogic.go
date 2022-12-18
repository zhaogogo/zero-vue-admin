package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRoleByUserIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserRoleByUserIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRoleByUserIDLogic {
	return &UserRoleByUserIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserRoleByUserIDLogic) UserRoleByUserID(in *pb.UserID) (*pb.UserRoleResponse, error) {
	userRoles, err := l.svcCtx.UserRoleModel.FindByUserID(l.ctx, l.svcCtx.Redis, in.ID)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return nil, err
		}
		return nil, errors.Wrap(err, "数据库查询失败")
	}
	pbUserRole := []*pb.UserRole{}
	for _, userRole := range userRoles {
		pbUserRole = append(pbUserRole, &pb.UserRole{ID: userRole.Id, UserID: userRole.UserId, RoleID: userRole.RoleId})
	}
	return &pb.UserRoleResponse{
		UserRoles: pbUserRole,
	}, nil
}
