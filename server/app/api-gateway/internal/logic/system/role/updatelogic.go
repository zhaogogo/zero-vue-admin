package role

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.RoleUpdateRequest) (resp *types.HttpCommonResponse, err error) {
	loginUser := l.ctx.Value("userName").(string)
	param := &systemservice.UpdateRoleRequest{
		ID:       req.ID,
		Role:     req.Role,
		Name:     req.Name,
		CreateBy: req.CreateBy,
		UpdateBy: loginUser,
		DeleteBy: req.DeleteBy,
	}
	_, err = l.svcCtx.SystemRpcClient.UpdateRole(l.ctx, param)
	if err != nil {
		return nil, errorx.NewByCode(err, errorx.GRPC_ERROR).WithMeta("SystemRpcClient.UpdateRole", err.Error(), param)
	}

	return &types.HttpCommonResponse{Code: 200, Msg: "OK"}, nil
}
