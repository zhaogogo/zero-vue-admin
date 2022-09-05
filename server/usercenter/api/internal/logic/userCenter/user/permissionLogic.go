package user

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/api/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/api/internal/types"
	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/rpc/system/systemuser"

	"github.com/zeromicro/go-zero/core/logx"
)

type PermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PermissionLogic {
	return &PermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PermissionLogic) Permission() (resp *types.Permission, err error) {
	userId, _ := l.ctx.Value("userId").(json.Number).Int64()
	userMenuTree, err := l.svcCtx.SystemUserRpc.UserPermission(l.ctx, &systemuser.UserPermissionRequest{UserId: userId})
	if err != nil {
		return nil, err
	}

	userinfo := types.User{}
	copier.Copy(&userinfo, userMenuTree.Userinfo)

	menus := l.getMenuTreeMap(userMenuTree.Menulists)
	resp = &types.Permission{
		UserInfo: userinfo,
		Menus:    menus,
	}
	return
}

func (l *PermissionLogic) getMenuTreeMap(menusList []*systemuser.MenuList) []*types.Menu {
	treeMap := make(map[int64][]*types.Menu)
	for _, menu := range menusList {
		treeMap[menu.ParentId] = append(treeMap[menu.ParentId], &types.Menu{
			ID:        menu.Id,
			ParentID:  menu.ParentId,
			Path:      menu.Path,
			Name:      menu.Name,
			Component: menu.Component,
			Hidden:    menu.Hidden,
			Meta: types.MenuMeta{
				Title: menu.Title,
				Icon:  menu.Icon,
			},
		})
	}
	menus := treeMap[0]
	for i := 0; i < len(menus); i++ {
		l.getChildrenList(menus[i], treeMap)
	}
	return menus
}

func (l *PermissionLogic) getChildrenList(menus *types.Menu, treeMap map[int64][]*types.Menu) {
	menus.Children = treeMap[menus.ID]
	for i := 0; i < len(menus.Children); i++ {
		l.getChildrenList(menus.Children[i], treeMap)
	}
}
