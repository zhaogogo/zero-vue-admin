package role

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAPIPermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAPIPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAPIPermissionLogic {
	return &UpdateAPIPermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAPIPermissionLogic) UpdateAPIPermission(req *types.UpdateRoleAPIPermissionRequest) (resp *types.HttpCommonResponse, err error) {
	param := &systemservice.UpdateCasbinPolicyRequest{V0: strconv.FormatUint(req.ID, 10)}
	for _, v := range req.CasbinRules {
		param.CasbinRules = append(param.CasbinRules, &systemservice.CasbinRule{
			V1: v.Api,
			V2: v.Method,
		})
	}
	_, err = l.svcCtx.SystemRpcClient.UpdateCasbinPolicy(l.ctx, param)
	if err != nil {
		return nil, errorx.NewByCode(err, errorx.GRPC_ERROR).WithMeta("SystemRpcClient.UpdateCasbinPolicy", err.Error(), param)
	}
	return &types.HttpCommonResponse{Code: 200, Msg: "OK"}, nil
}
