package role

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMenuPermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMenuPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuPermissionLogic {
	return &UpdateMenuPermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMenuPermissionLogic) UpdateMenuPermission(req *types.UpdateRoleMenuPermissionRequest) (resp *types.HttpCommonResponse, err error) {
	param := &systemservice.UpdateRoleMenusRequest{RoleID: req.ID, MenuIDList: req.MenuIDList}
	_, err = l.svcCtx.SystemRpcClient.UpdateRoleMenus(l.ctx, param)
	if err != nil {
		return nil, errorx.NewByCode(err, errorx.GRPC_ERROR).WithMeta("SystemRpcClient.UpdateRoleMenus", err.Error(), param)
	}
	return &types.HttpCommonResponse{Code: 200, Msg: "OK"}, nil
}
