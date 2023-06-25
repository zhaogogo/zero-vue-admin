package menu

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"

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

func (l *UpdateLogic) Update(req *types.MenuUpdateRequest) (resp *types.HttpCommonResponse, err error) {
	param := &systemservice.UpdateMenuRequest{
		Id:        req.ID,
		ParentID:  req.ParentID,
		Name:      req.Name,
		Path:      req.Path,
		Component: req.Component,
		Title:     req.Meta.Title,
		Icon:      req.Meta.Icon,
		Sort:      req.Sort,
	}
	if req.Hidden {
		param.Hiddent = 1
	} else {
		param.Hiddent = 0
	}
	_, err = l.svcCtx.SystemRpcClient.UpdateMenu(l.ctx, param)
	if err != nil {
		return nil, errorx.New(err, "更新菜单失败").WithMeta("SystemRpcClient.UpdateMenu", err.Error(), param)
	}
	return &types.HttpCommonResponse{Code: 200, Msg: "OK"}, nil
}
