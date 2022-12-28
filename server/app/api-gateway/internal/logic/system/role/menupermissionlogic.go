package role

import (
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"
	"google.golang.org/grpc/status"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuPermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuPermissionLogic {
	return &MenuPermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuPermissionLogic) MenuPermission(req *types.RoleMenuPermissionRequest) (resp *types.RoleMenuPermissionResponse, err error) {
	msgErrList := errorx.MsgErrList{}
	menus := []types.MenuResp{}
	roleMenuParam := &systemservice.RoleID{ID: req.ID}
	res, err := l.svcCtx.SystemRpcClient.RoleMenuByRoleID(l.ctx, roleMenuParam)
	if err != nil {
		s, _ := status.FromError(err)
		if s.Message() == sql.ErrNoRows.Error() {
			res = new(systemservice.RoleMenuResponse)
		} else {
			return nil, errorx.NewByCode(err, errorx.GRPC_ERROR).WithMeta("SystemRpcClient.RoleMenuByRoleID", err.Error(), roleMenuParam)
		}
	}
	mr.MapReduce(
		func(source chan<- interface{}) {
			for _, v := range res.Rolemenus {
				source <- v.MenuID
			}
		},
		func(item interface{}, writer mr.Writer, cancel func(error)) {
			menuid := item.(uint64)
			menuDetailParam := &systemservice.MenuID{ID: menuid}
			menu, err := l.svcCtx.SystemRpcClient.MenuDetail(l.ctx, menuDetailParam)
			if err != nil {
				msgErrList.WithMeta("SystemRpcClient.MenuDetail", err.Error(), menuDetailParam)
			} else {
				if menu.DeleteTime == 0 {
					writer.Write(menu)
				}
			}
		},
		func(pipe <-chan interface{}, writer mr.Writer, cancel func(error)) {
			for v := range pipe {
				menu := v.(*systemservice.Menu)
				menus = append(menus, types.MenuResp{
					ID:        menu.ID,
					ParentId:  menu.ParentID,
					Name:      menu.Name,
					Path:      menu.Path,
					Component: menu.Component,
					Sort:      menu.Sort,
					Hidden:    menu.Hiddent == 1,
					MenuMeta:  types.MenuMeta{Title: menu.Title, Icon: menu.Icon},
				})
			}
		},
	)

	return &types.RoleMenuPermissionResponse{
		HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: "OK"},
		List:               menus,
	}, nil
}
