package menu

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"
	"google.golang.org/grpc/status"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.MenuDetailRequest) (resp *types.MenuDetailResponse, err error) {
	msgErrList := errorx.MsgErrList{}

	param := &systemservice.MenuID{ID: req.ID}
	pmenu, err := l.svcCtx.SystemRpcClient.MenuDetail(l.ctx, param)
	if err != nil {
		s, _ := status.FromError(err)
		if s.Message() == sql.ErrNoRows.Error() {
			return nil, errorx.New(err, fmt.Sprintf("获取菜单失败, id=%v无数据", req.ID)).WithMeta("SystemRpcClient.MenuDetail", err.Error(), param)
		}
		return nil, errorx.New(err, "获取菜单失败").WithMeta("SystemRpcClient.MenuDetail", err.Error(), param)
	}

	userMenuParamsParam := &systemservice.Empty{}
	usermenuparams, err := l.svcCtx.SystemRpcClient.UserAllMenuParams(l.ctx, userMenuParamsParam)
	if err != nil {
		s, _ := status.FromError(err)
		if s.Message() == sql.ErrNoRows.Error() {
			usermenuparams = new(systemservice.UserMenuParamsResponse)
		} else {
			msgErrList.WithMeta("SystemRpcClient.UserAllMenuParams", err.Error(), userMenuParamsParam)
		}
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

	menuParam := []types.Parameter{}
	for _, usermenuparam := range usermenuparams.UserMenuParams {
		if menu.ID == usermenuparam.MenuID {
			menuParam = append(menuParam, types.Parameter{
				ID:     usermenuparam.ID,
				UserID: usermenuparam.UserID,
				MenuID: usermenuparam.MenuID,
				Type:   usermenuparam.Type,
				Key:    usermenuparam.Key,
				Value:  usermenuparam.Value,
			})
		}
	}
	menu.Parameters = menuParam

	var (
		msg      = "OK"
		errcount = len(msgErrList.List)
	)
	if errcount != 0 {
		msg = fmt.Sprintf("Not OK(%v) %v", errcount, msgErrList.List)
	}

	return &types.MenuDetailResponse{
		HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: msg, Meta: msgErrList.List},
		Detail:             menu,
	}, nil
}
