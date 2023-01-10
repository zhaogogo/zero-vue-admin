package logic

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/model"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateRoleLogic) UpdateRole(in *pb.UpdateRoleRequest) (*pb.Empty, error) {
	err := l.svcCtx.RoleModel.Update(l.ctx, &model.Role{
		Id:         in.ID,
		Role:       in.Role,
		Name:       in.Name,
		CreateBy:   in.CreateBy,
		UpdateBy:   in.UpdateBy,
		DeleteBy:   in.DeleteBy,
		DeleteTime: sql.NullTime{},
	})
	if err != nil {
		return nil, errors.Wrap(err, "数据库错误")
	}

	return &pb.Empty{}, nil
}
