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

type CreateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoleLogic {
	return &CreateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateRoleLogic) CreateRole(in *pb.CreateRoleRequest) (*pb.Empty, error) {
	_, err := l.svcCtx.RoleModel.Insert(l.ctx, &model.Role{
		Id:         0,
		Role:       in.Role,
		Name:       in.Name,
		CreateBy:   in.CreateBy,
		UpdateBy:   in.CreateBy,
		DeleteBy:   "",
		DeleteTime: sql.NullTime{},
	})
	if err != nil {
		return nil, errors.Wrap(err, "数据库错误")
	}
	return &pb.Empty{}, nil
}
