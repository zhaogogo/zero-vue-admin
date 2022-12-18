package menu

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMenuLogic {
	return &AddMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddMenuLogic) AddMenu(req *types.AddMenuRequest) (resp *types.HttpCommonResponse, err error) {
	param := &systemservice.AddMenuRequest{
		ParentID:  req.ParentID,
		Name:      req.Name,
		Path:      req.Name,
		Component: req.Component,
		Title:     req.Title,
		Icon:      req.Icon,
		Sort:      req.Sort,
	}
	if req.Hidden {
		param.Hiddent = 1
	}
	_, err = l.svcCtx.SystemRpcClient.AddMenu(l.ctx, param)
	if err != nil {
		return nil, errorx.NewByCode(err, errorx.GRPC_ERROR).WithMeta("SystemRpcClient.AddMenu", err.Error(), param)
	}
	return &types.HttpCommonResponse{Code: 200, Msg: "OK"}, nil
}
