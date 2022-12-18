package menu

import (
	"context"
	"database/sql"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"
	"google.golang.org/grpc/status"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuInfoLogic {
	return &MenuInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuInfoLogic) MenuInfo(req *types.MenuInfoRequest) (resp *types.MenuInfoResponse, err error) {
	param := &systemservice.MenuID{ID: req.ID}
	pmenu, err := l.svcCtx.SystemRpcClient.MenuInfo(l.ctx, param)
	if err != nil {
		s, _ := status.FromError(err)
		if s.Message() == sql.ErrNoRows.Error() {
			return nil, errorx.NewByCode(err, errorx.DB_NOTFOUND).WithMeta("SystemRpcClient.MenuInfo", err.Error(), param)
		}
		return nil, errorx.NewByCode(err, errorx.GRPC_ERROR).WithMeta("SystemRpcClient.MenuInfo", err.Error(), param)
	}

	menu := types.Menu{
		ID:        pmenu.ID,
		ParentId:  pmenu.ParentID,
		Name:      pmenu.Name,
		Path:      pmenu.Path,
		Component: pmenu.Component,
		Sort:      pmenu.Sort,
		Hidden:    !(pmenu.Hiddent == 0),
		MenuMeta:  types.MenuMeta{Title: pmenu.Title, Icon: pmenu.Icon},
	}

	return &types.MenuInfoResponse{
		HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: "OK"},
		MenuInfo:           menu,
	}, nil
}
