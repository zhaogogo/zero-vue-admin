package user

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"
	"google.golang.org/grpc/status"
	"sync"

	"github.com/zeromicro/go-zero/core/logx"
)

type CurrentUserSetingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCurrentUserSetingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CurrentUserSetingLogic {
	return &CurrentUserSetingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CurrentUserSetingLogic) CurrentUserSeting() (resp *types.CurrentUserSetResponse, err error) {
	msgErrList := errorx.MsgErrList{}
	roleInfo := make(map[uint64]*systemservice.Role)
	userinfo := l.ctx.Value("userinfo").(*systemservice.User)
	userroleinfo := l.ctx.Value("userroleinfo").([]*systemservice.UserRole)
	var (
		userpageset *systemservice.UserPageSetResponse
		wg          sync.WaitGroup
	)
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			if e := recover(); e != nil {
				logx.Error(e)
			}
		}()
		userPageSetParam := &systemservice.UserID{ID: userinfo.ID}
		userpageset, err = l.svcCtx.SystemRpcClient.UserPageSet(l.ctx, userPageSetParam)
		if err != nil {
			s, _ := status.FromError(err)
			if s.Message() == sql.ErrNoRows.Error() {
				userpageset = new(systemservice.UserPageSetResponse)
				userpageset.ID = 0
				userpageset.UserId = userinfo.ID
				userpageset.Avatar = ""
				userpageset.DefaultRouter = "dashboard"
				userpageset.SideMode = "#191a23"
				userpageset.ActiveTextColor = "#1890ff"
				userpageset.TextColor = "#fff"
			} else {
				msgErrList.WithMeta("SystemRpcClient.UserPageSet", err.Error(), userPageSetParam)
			}
		}
	}()

	mr.MapReduce(
		func(source chan<- interface{}) {
			for _, v := range userroleinfo {
				source <- v.RoleID
			}
		},
		func(item interface{}, writer mr.Writer, cancel func(error)) {
			roleid := item.(uint64)
			roleDetailParam := &systemservice.RoleID{ID: roleid}
			role, err := l.svcCtx.SystemRpcClient.RoleDetail(l.ctx, roleDetailParam)
			if err != nil {
				msgErrList.WithMeta("SystemRpcClient.RoleDetail", err.Error(), roleDetailParam)
				return
			}
			writer.Write(role)
		},
		func(pipe <-chan interface{}, writer mr.Writer, cancel func(error)) {
			for v := range pipe {
				role := v.(*systemservice.Role)
				roleInfo[role.ID] = role
			}
		},
	)

	roles := []types.Role{}
	currentRole := types.Role{}
	for roleid, role := range roleInfo {
		if userinfo.CurrentRole == roleid {
			currentRole = types.Role{
				ID:         role.ID,
				Role:       role.Role,
				Name:       role.Name,
				CreateBy:   role.CreateBy,
				CreateTime: role.CreateTime,
				UpdateBy:   role.UpdateBy,
				UpdateTime: role.UpdateTime,
				DeleteBy:   role.DeleteBy,
				DeleteTime: role.DeleteTime,
			}
		}
		roles = append(roles, types.Role{
			ID:         role.ID,
			Role:       role.Role,
			Name:       role.Name,
			CreateBy:   role.CreateBy,
			CreateTime: role.CreateTime,
			UpdateBy:   role.UpdateBy,
			UpdateTime: role.UpdateTime,
			DeleteBy:   role.DeleteBy,
			DeleteTime: role.DeleteTime,
		})
	}
	wg.Wait()
	var msg string = "OK"
	if len(msgErrList.List) != 0 {
		msg = fmt.Sprintf("Not OK(%d), %v", len(msgErrList.List), msgErrList.List)
	}
	return &types.CurrentUserSetResponse{
		HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: msg, Meta: msgErrList.List},
		User: types.User{
			ID:         userinfo.ID,
			Name:       userinfo.Name,
			NickName:   userinfo.NickName,
			UserType:   userinfo.UserType,
			Email:      userinfo.Email,
			Phone:      userinfo.Phone,
			Department: userinfo.Department,
			Position:   userinfo.Position,
			CreateBy:   userinfo.CreateBy,
			CreateTime: userinfo.CreateTime,
			UpdateBy:   userinfo.UpdateBy,
			UpdateTime: userinfo.UpdateTime,
			DeleteBy:   userinfo.DeleteBy,
			DeleteTime: userinfo.DeleteTime,
		},
		UserPageSet: types.UserPageSet{
			Name:            userinfo.Name,
			NickName:        userinfo.NickName,
			Avatar:          userpageset.Avatar,
			DefaultRouter:   userpageset.DefaultRouter,
			SideMode:        userpageset.SideMode,
			ActiveTextColor: userpageset.ActiveTextColor,
			TextColor:       userpageset.TextColor,
		},
		Roles:       roles,
		CurrentRole: currentRole,
	}, nil
}
