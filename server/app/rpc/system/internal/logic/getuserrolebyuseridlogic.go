package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserRoleByUserIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserRoleByUserIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserRoleByUserIDLogic {
	return &GetUserRoleByUserIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserRoleByUserIDLogic) GetUserRoleByUserID(in *pb.UserID) (*pb.UserRoleList, error) {
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
	return &pb.UserRoleList{
		UserRole: pbUserRole,
	}, nil
}
