package menu

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/responseerror/errorx"
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

func (l *CreateLogic) Create(req *types.MenuCreateRequest) (resp *types.HttpCommonResponse, err error) {
	param := &systemservice.CreateMenuRequest{
		ParentID:  req.ParentID,
		Name:      req.Name,
		Path:      req.Name,
		Component: req.Component,
		Title:     req.Meta.Title,
		Icon:      req.Meta.Icon,
		Sort:      req.Sort,
	}
	if req.Hidden {
		param.Hiddent = 1
	}
	_, err = l.svcCtx.SystemRpcClient.CreateMenu(l.ctx, param)
	if err != nil {
		return nil, errorx.New(err, "创建菜单失败").WithMeta("SystemRpcClient.CreateMenu", err.Error(), param)
	}
	return &types.HttpCommonResponse{Code: 200, Msg: "OK"}, nil
}
