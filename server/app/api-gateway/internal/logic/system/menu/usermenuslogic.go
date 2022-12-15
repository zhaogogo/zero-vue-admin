package menu

import (
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/utils"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"
	"google.golang.org/grpc/status"
	"sort"
	"strings"
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
		msgErrList        = errorx.MsgErrList{}
		msg        string = "OK"
	)
	userid, err := utils.GetUserIdWithJWT(l.ctx)
	if err != nil {
		return nil, errorx.New(err, errorx.SERVER_COMMON_ERROR, "JWT转换用户ID类型失败")
	}

	userroles, err := l.svcCtx.SystemRpcClient.GetUserRoleByUserID(l.ctx, &systemservice.UserID{ID: userid})
	if err != nil {
		return nil, err
	}

	// 获取menuid列表
	mr.MapReduce(
		func(source chan<- interface{}) {
			for _, userrole := range userroles.UserRole {
				source <- userrole.RoleID
			}
		},
		func(item interface{}, writer mr.Writer, cancel func(error)) {
			roleid := item.(uint64)
			roleMenu, err := l.svcCtx.SystemRpcClient.GetRoleMenuByRoleID(l.ctx, &systemservice.RoleID{ID: roleid})
			if err != nil {
				l.Error(err)
				msgErrList.Append(err.Error())
			} else {
				writer.Write(roleMenu)
			}
		},
		func(pipe <-chan interface{}, writer mr.Writer, cancel func(error)) {
			for v := range pipe {
				rolemenu, ok := v.(*systemservice.RoleMenuList)
				if ok {
					for _, rm := range rolemenu.Rolemenu {
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
			menuInfo, err := l.svcCtx.SystemRpcClient.MenuInfo(l.ctx, &systemservice.MenuID{ID: menuid})
			if err != nil {
				l.Error(err)
				msgErrList.Append(err.Error())
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
					MenuMeta:  types.MenuMeta{Icon: menuInfo.Icon, Title: menuInfo.Title},
				})
			}
		},
	)
	// 获取参数
	params := []types.Parameter{}
	userMenuParams, err := l.svcCtx.SystemRpcClient.GetUserMenuParams(l.ctx, &systemservice.UserID{ID: userid})
	s, _ := status.FromError(err)
	if s.Message() == sql.ErrNoRows.Error() {

	} else if err != nil {
		msgErrList.Append(err.Error())
	} else {
		for _, p := range userMenuParams.UserMenuParams {
			params = append(params, types.Parameter{
				UserID: p.UserID,
				MenuID: p.MenuID,
				Type:   p.Type,
				Key:    p.Key,
				Value:  p.Value,
			})
		}
	}

	menuTree := genMenuTreeMap(menus, params)

	if len(msgErrList.List) != 0 {
		msg = strings.Join(msgErrList.List, " | ")
	}
	return &types.UserMenuResponse{
		HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: msg},
		Menus:              menuTree,
	}, nil
}

func genMenuTreeMap(menus []types.Menu, params []types.Parameter) []types.Menu {
	sort.Slice(menus, func(i, j int) bool {
		return menus[i].Sort < menus[j].Sort
	})
	var treeMap = make(map[uint64][]types.Menu)
	for _, menu := range menus {
		for _, param := range params {
			if param.MenuID == menu.ID {
				menu.Parameters = append(menu.Parameters, types.Parameter{UserID: param.UserID, MenuID: param.MenuID, Type: param.Type, Key: param.Key, Value: param.Value})
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
