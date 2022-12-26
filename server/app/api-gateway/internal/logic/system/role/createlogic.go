package role

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.RoleCreateRequest) (resp *types.HttpCommonResponse, err error) {
	loginUser := l.ctx.Value("userName").(string)
	param := &systemservice.CreateRoleRequest{
		Role:     req.Role,
		Name:     req.Name,
		CreateBy: loginUser,
	}
	_, err = l.svcCtx.SystemRpcClient.CreateRole(l.ctx, param)
	if err != nil {
		return nil, errorx.NewByCode(err, errorx.GRPC_ERROR).WithMeta("SystemRpcClient.CreateRole", err.Error(), param)
	}

	return &types.HttpCommonResponse{Code: 200, Msg: "OK"}, nil
}
