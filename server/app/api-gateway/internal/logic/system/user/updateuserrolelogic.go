package user

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserRoleLogic {
	return &UpdateUserRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserRoleLogic) UpdateUserRole(req *types.UpdateUserRoleRequest) (resp *types.HttpCommonResponse, err error) {
	updateUserRoleParam := &systemservice.UpdateUserRoleRequest{UserID: req.ID, RoleList: req.RoleList}
	_, err = l.svcCtx.SystemRpcClient.UpdateUserRole(l.ctx, updateUserRoleParam)
	if err != nil {
		return nil, errorx.New(err, "更新用户角色失败").WithMeta("*SystemRpcClient.UpdateUserRole", err.Error(), updateUserRoleParam)
	}
	return &types.HttpCommonResponse{Code: 200, Msg: "OK"}, nil
}
