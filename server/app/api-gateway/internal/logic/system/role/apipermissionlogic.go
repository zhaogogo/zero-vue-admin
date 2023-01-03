package role

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type APIPermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAPIPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *APIPermissionLogic {
	return &APIPermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *APIPermissionLogic) APIPermission(req *types.RoleAPIPermissionRequest) (resp *types.RoleAPIPermissionResponse, err error) {
	param := &systemservice.RoleID{ID: req.ID}
	res, err := l.svcCtx.SystemRpcClient.CasbinPolicyByRoleID(l.ctx, param)
	if err != nil {
		return nil, errorx.NewByCode(err, errorx.GRPC_ERROR).WithMeta("SystemRpcClient.CasbinPolicyByRoleID", err.Error(), param)
	}
	policy := []types.CasbinPolicy{}
	for _, p := range res.Policy {
		policy = append(policy, types.CasbinPolicy{API: p.Api, Method: p.Method})
	}
	return &types.RoleAPIPermissionResponse{
		HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: "OK"},
		Policy:             policy,
	}, nil
}
