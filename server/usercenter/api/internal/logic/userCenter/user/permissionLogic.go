package user

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/api/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PermissionLogic {
	return &PermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PermissionLogic) Permission() (resp *types.Permission, err error) {
	// todo: add your logic here and delete this line

	return
}
