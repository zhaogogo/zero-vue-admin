package api

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

func (l *UpdateLogic) Update(req *types.APIUpdateRequest) (resp *types.HttpCommonResponse, err error) {
	param := &systemservice.API{
		ID:       req.ID,
		API:      req.Api,
		Group:    req.Group,
		Describe: req.Describe,
		Method:   req.Method,
	}
	_, err = l.svcCtx.SystemRpcClient.UpdateAPI(l.ctx, param)
	if err != nil {
		return nil, errorx.NewByCode(err, errorx.GRPC_ERROR).WithMeta("SystemRpcClient.UpdateAPI", err.Error(), param)
	}
	return &types.HttpCommonResponse{
		Code: 200,
		Msg:  "OK",
	}, nil
}
