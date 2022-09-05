package logic

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/rpc/system/pb"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserPermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserPermissionLogic {
	return &UserPermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserPermissionLogic) UserPermission(in *pb.UserPermissionRequest) (*pb.UserPermissionResponse, error) {
	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, errors.Wrap(err, fmt.Sprintf("用户不存在, userId: %v", in.UserId))
		}
		return nil, errors.Wrap(err, "数据库错误")
	}
	fmt.Printf("@@@ %#v", userInfo)
	userRoles, err := l.svcCtx.UserRoleModel.FindByUserID(l.ctx, userInfo.Id)
	if err != nil {
		return nil, errors.Wrap(err, "查询用户角色失败")
	}
	roles := []string{}
	for _, userRole := range userRoles {
		roles = append(roles, strconv.Itoa(int(userRole.RoleId)))
	}
	menus, err := l.svcCtx.MenuModel.FindMenusByRoles(l.ctx, roles...)
	if err != nil {
		return nil, errors.Wrap(err, "查询用户Menu失败")
	}

	pbUser := &pb.User{}
	pbMenuLists := []*pb.MenuList{}
	copier.Copy(pbUser, userInfo)
	pbUser.CreateAt = userInfo.CreateAt.Unix()
	pbUser.UpdateAt = userInfo.UpdateAt.Unix()

	for _, menu := range menus {
		menuList := pb.MenuList{}
		copier.Copy(&menuList, *menu)
		pbMenuLists = append(pbMenuLists, &menuList)
	}
	return &pb.UserPermissionResponse{
		Userinfo:  pbUser,
		Menulists: pbMenuLists,
	}, nil
}
