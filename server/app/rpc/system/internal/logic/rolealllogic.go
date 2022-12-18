package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleAllLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleAllLogic {
	return &RoleAllLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleAllLogic) RoleAll(in *pb.Empty) (*pb.RoleAllResponse, error) {
	mRole, err := l.svcCtx.RoleModel.FindAll_NC(l.ctx)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return nil, err
		}
		return nil, errors.Wrapf(err, "查询数据库错误, 表: role 字段:AllRole")
	}

	list := []*pb.Role{}
	for _, role := range mRole {
		pr := &pb.Role{
			ID:         role.Id,
			Role:       role.Role,
			Name:       role.Name,
			CreateBy:   role.CreateBy,
			CreateTime: role.CreateTime.Unix(),
			UpdateBy:   role.UpdateBy,
			UpdateTime: role.UpdateTime.Unix(),
			DeleteBy:   role.DeleteBy,
		}
		if role.DeleteTime.Valid {
			pr.DeleteTime = role.DeleteTime.Time.Unix()
		} else {
			pr.DeleteTime = 0
		}
		list = append(list, pr)
	}

	return &pb.RoleAllResponse{
		Roles: list,
	}, nil
}
