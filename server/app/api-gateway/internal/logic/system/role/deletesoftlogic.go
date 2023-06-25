package role

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSoftLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSoftLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSoftLogic {
	return &DeleteSoftLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSoftLogic) DeleteSoft(req *types.RoleDeleteSoftRequest) (resp *types.HttpCommonResponse, err error) {
	username := l.ctx.Value("userName").(string)
	param := &systemservice.DeleteSoftRoleRequest{RoleID: req.ID, DeleteBy: username, State: req.State}
	_, err = l.svcCtx.SystemRpcClient.DeleteSoftRole(l.ctx, param)
	if err != nil {
		return nil, errorx.New(err, "禁用角色失败").WithMeta("*SystemRpcClient.DeleteSoftRole", err.Error(), param)
	}
	return &types.HttpCommonResponse{Code: 200, Msg: "OK"}, nil
}
