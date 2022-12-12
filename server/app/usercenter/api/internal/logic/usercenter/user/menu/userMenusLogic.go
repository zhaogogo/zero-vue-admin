package menu

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	errorx2 "github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/usercenter/api/internal/pkg/logiccommon"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/usercenter/api/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/usercenter/api/internal/types"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/usercenter/model"
	"sort"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *UserMenusLogic) UserMenus() (resp *types.MenusResponse, err error) {
	userid, err := logiccommon.GetUserIdWithJWT(l.ctx)
	if err != nil {
		return nil, errorx2.New(errorx2.SERVER_COMMON_ERROR, "获取用户ID失败")
	}
	menus, err := l.getmenus(userid)
	if err != nil {
		logx.Error(err)
		return nil, errorx2.New(errorx2.SERVER_COMMON_ERROR, "获取menu数据失败")
	}
	params, err := l.svcCtx.UserMenuParamsModel.FindByUserID(l.ctx, userid)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, errorx2.New(errorx2.SERVER_COMMON_ERROR, "获取user parametes为空")
		}
		return nil, errors.Wrap(err, "获取user parametes为空")
	}
	menutree := genMenuTreeMap(menus, params)

	return &types.MenusResponse{
		HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: "ok"},
		Menus:              menutree,
	}, nil
}
func genMenuTreeMap(menus []model.Menu, params []model.UserMenuParams) []types.Menu {
	sort.Slice(menus, func(i, j int) bool {
		return menus[i].Sort < menus[j].Sort
	})
	var treeMap = make(map[uint64][]types.Menu)
	for _, menu := range menus {
		m := types.Menu{
			ID:        menu.Id,
			ParentId:  menu.ParentId,
			Name:      menu.Name,
			Path:      menu.Path,
			Sort:      menu.Sort,
			Component: menu.Component,
			MenuMeta:  types.MenuMeta{Title: menu.Title, Icon: menu.Icon},
		}
		for _, param := range params {
			if param.MenuId == m.ID {
				m.Parameters = append(m.Parameters, types.Parameter{UserID: param.UserId, Type: param.Type, Key: param.Key, Value: param.Value})
			}
		}
		treeMap[menu.ParentId] = append(treeMap[menu.ParentId], m)
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

	return
}

func (l *UserMenusLogic) getmenus(userid uint64) ([]model.Menu, error) {
	roleIDs, err := l.svcCtx.UserRoleModel.FindByUserID(l.ctx, userid)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, errorx2.New(errorx2.DB_NOTFOUND, "获取角色ID为空")
		}
		return nil, errors.Wrap(err, "获取角色id失败")
	}

	menuIDs, err := l.svcCtx.RoleMenuModel.FindByRoleID(l.ctx, roleIDs)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, errorx2.New(errorx2.SERVER_COMMON_ERROR, "获取menuIDs为空")
		}
		return nil, errors.Wrap(err, "获取menuIDs失败")
	}
	//fmt.Printf("===> menu_id: %v\n", menuIDs)
	menus, err := l.svcCtx.MenuModel.FindByIDs(l.ctx, menuIDs)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, errorx2.New(errorx2.SERVER_COMMON_ERROR, "获取menu为空")
		}
		return nil, errors.Wrap(err, "获取menus数据失败")
	}

	return menus, nil
}
