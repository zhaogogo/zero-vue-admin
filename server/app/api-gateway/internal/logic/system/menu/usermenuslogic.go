package menu

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"
	"google.golang.org/grpc/status"
	"sort"
)

type UserMenusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserMenusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserMenusLogic {
	return &UserMenusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserMenusLogic) UserMenus() (resp *types.UserMenuResponse, err error) {
	var (
		setMenus   = make(map[uint64]int)
		menus      []types.Menu
		msgErrList = errorx.MsgErrList{}
	)
	hasInRole := false
	userid := l.ctx.Value("user_id").(uint64)
	userinfo := l.ctx.Value("userinfo").(*systemservice.User)
	userroles := l.ctx.Value("userroleinfo").([]*systemservice.UserRole)
	for _, v := range userroles {
		if v.RoleID == userinfo.CurrentRole {
			hasInRole = true
			break
		}
	}
	if !hasInRole {
		return &types.UserMenuResponse{HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: "OK", Meta: msgErrList.List}, Menus: nil}, nil
	}

	// 获取menuid列表
	mr.MapReduce(
		func(source chan<- interface{}) {
			//for _, userrole := range userroles.UserRoles {
			//	source <- userrole.RoleID
			//}
			source <- userinfo.CurrentRole
		},
		func(item interface{}, writer mr.Writer, cancel func(error)) {
			roleid := item.(uint64)
			getRoleMenuByRoleIDParam := &systemservice.RoleID{ID: roleid}
			roleMenu, err := l.svcCtx.SystemRpcClient.RoleMenuByRoleID(l.ctx, getRoleMenuByRoleIDParam)
			if err != nil {
				l.Error(err)
				msgErrList.WithMeta("SystemRpcClient.RoleMenuByRoleID", err.Error(), getRoleMenuByRoleIDParam)
			} else {
				writer.Write(roleMenu)
			}
		},
		func(pipe <-chan interface{}, writer mr.Writer, cancel func(error)) {
			for v := range pipe {
				rolemenu, ok := v.(*systemservice.RoleMenuResponse)
				if ok {
					for _, rm := range rolemenu.Rolemenus {
						setMenus[rm.MenuID] += 1
					}
				} else {
					logx.Errorf("mr reducer断言失败, 实际类型: (%T), 断言类型: (*systemservice.RoleMenuList)", v, v)
				}

			}
		},
	)
	//动态获取menus，不返回软删除menus
	mr.MapReduce(
		func(source chan<- interface{}) {
			for menuId, _ := range setMenus {
				source <- menuId
			}
		},
		func(item interface{}, writer mr.Writer, cancel func(error)) {
			menuid := item.(uint64)
			menuInfoParam := &systemservice.MenuID{ID: menuid}
			menuInfo, err := l.svcCtx.SystemRpcClient.MenuDetail(l.ctx, menuInfoParam)
			if err != nil {
				l.Error(err)
				msgErrList.WithMeta("SystemRpcClient.MenuDetail", err.Error(), menuInfoParam)
			} else {
				if menuInfo.DeleteTime == 0 {
					writer.Write(menuInfo)
				}
			}
		},
		func(pipe <-chan interface{}, writer mr.Writer, cancel func(error)) {
			for menuinfo := range pipe {
				menuInfo := menuinfo.(*systemservice.Menu)
				menus = append(menus, types.Menu{
					ID:        menuInfo.ID,
					ParentId:  menuInfo.ParentID,
					Name:      menuInfo.Name,
					Path:      menuInfo.Path,
					Component: menuInfo.Component,
					Sort:      menuInfo.Sort,
					Hidden:    menuInfo.Hiddent == 1,
					MenuMeta:  types.MenuMeta{Icon: menuInfo.Icon, Title: menuInfo.Title},
				})
			}
		},
	)
	// 获取参数
	params := []types.Parameter{}
	getUserMenuParamsParam := &systemservice.UserID{ID: userid}
	userMenuParams, err := l.svcCtx.SystemRpcClient.UserMenuParamsByUserID(l.ctx, getUserMenuParamsParam)
	s, _ := status.FromError(err)
	if s.Message() == sql.ErrNoRows.Error() {
		msgErrList.WithMeta("SystemRpcClient.GetUserMenuParams", err.Error(), getUserMenuParamsParam)
	} else if err != nil {
		msgErrList.WithMeta("SystemRpcClient.GetUserMenuParams", err.Error(), getUserMenuParamsParam)
	} else {
		for _, p := range userMenuParams.UserMenuParams {
			params = append(params, types.Parameter{
				ID:     p.ID,
				UserID: p.UserID,
				MenuID: p.MenuID,
				Type:   p.Type,
				Key:    p.Key,
				Value:  p.Value,
			})
		}
	}

	menuTree := genMenuTreeMap(menus, params)

	var (
		msg   string = "OK"
		count int    = len(msgErrList.List)
	)

	if count != 0 {
		msg = fmt.Sprintf("Not OK(%d) %v", count, msgErrList.List)
	}
	return &types.UserMenuResponse{
		HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: msg, Meta: msgErrList.List},
		Menus:              menuTree,
	}, nil
}

func genMenuTreeMap(menus []types.Menu, params []types.Parameter) []types.Menu {
	sort.Slice(menus, func(i, j int) bool {
		return menus[i].Sort < menus[j].Sort
	})
	sort.Slice(params, func(i, j int) bool {
		return params[i].UserID < params[j].UserID
	})
	var treeMap = make(map[uint64][]types.Menu)
	for _, menu := range menus {
		for _, param := range params {
			if param.MenuID == menu.ID {
				menu.Parameters = append(menu.Parameters, types.Parameter{ID: param.ID, UserID: param.UserID, MenuID: param.MenuID, Type: param.Type, Key: param.Key, Value: param.Value})
			}
		}
		treeMap[menu.ParentId] = append(treeMap[menu.ParentId], menu)
	}
	ms := []types.Menu{}

	for _, m := range treeMap[0] {
		ms = append(ms, m)
	}

	for i := 0; i < len(ms); i++ {
		mergChildren(&ms[i], treeMap)
	}
	return ms
}

func mergChildren(menu *types.Menu, treemap map[uint64][]types.Menu) {
	for _, m := range treemap[menu.ID] {
		menu.Children = append(menu.Children, m)
	}
	for i := 0; i < len(menu.Children); i++ {
		mergChildren(&menu.Children[i], treemap)
	}
}
