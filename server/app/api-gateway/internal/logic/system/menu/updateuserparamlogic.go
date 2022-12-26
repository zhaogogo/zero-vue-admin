package menu

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserParamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserParamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserParamLogic {
	return &UpdateUserParamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserParamLogic) UpdateUserParam(req *types.MenuUserParamRequest) (resp *types.HttpCommonResponse, err error) {
	param := &systemservice.UpdateUserMenuParamsRequest{
		MenuId:         req.ID,
		UserMenuParams: []*systemservice.UserMenuParams{},
	}
	for _, v := range req.Parameters {
		param.UserMenuParams = append(param.UserMenuParams, &systemservice.UserMenuParams{
			ID:     v.ID,
			UserID: v.UserID,
			MenuID: req.ID,
			Type:   v.Type,
			Key:    v.Key,
			Value:  v.Value,
		})
	}

	_, err = l.svcCtx.SystemRpcClient.UpdateUserMenuParams(l.ctx, param)
	if err != nil {
		return nil, errorx.NewByCode(err, errorx.GRPC_ERROR).WithMeta("SystemRpcClient.UpdateUserMenuParams", err.Error(), param)
	}
	return &types.HttpCommonResponse{Code: 200, Msg: "OK"}, nil
}
