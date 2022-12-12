package menu

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	errorx2 "github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/usercenter/api/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/usercenter/api/internal/types"
)

type AllMenusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAllMenusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AllMenusLogic {
	return &AllMenusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AllMenusLogic) AllMenus() (resp *types.AllMenusResponse, err error) {
	var msg = "ok"
	menus, err := l.svcCtx.MenuModel.FindAll(l.ctx)
	if err != nil {
		return nil, errorx2.New(errorx2.SERVER_COMMON_ERROR, "获取menu全量数据失败")
	}

	params, err := l.svcCtx.UserMenuParamsModel.FindAll(l.ctx)
	if err != nil {
		//if errors.Is(err, sqlx.ErrNotFound) {
		//	msg = "获取user parametes为空" + err.Error()
		//}
		logx.Errorf("获取user parametes为空, error: %v", err)
		msg = "获取user parametes为空" + err.Error()
	}

	menutree := genMenuTreeMap(menus, params)

	return &types.AllMenusResponse{
		HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: msg},
		Total:              uint64(len(menus)),
		List:               menutree,
	}, nil
}
