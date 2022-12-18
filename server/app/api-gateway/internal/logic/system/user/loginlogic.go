package user

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/status"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	loginParam := &systemservice.LoginRequest{Name: req.UserName, PassWord: req.PassWord}
	res, err := l.svcCtx.SystemRpcClient.Login(l.ctx, loginParam)
	if err != nil {
		return nil, errorx.NewByCode(err, errorx.USERPASSWORDERROR).WithMeta("*SystemRpcClient.Login", err.Error(), loginParam)
	}

	var (
		userpageset *systemservice.UserPageSet
		user        *systemservice.User
		rolelist    = []string{}
		msgErrList  = errorx.MsgErrList{}
	)

	gCtx, _ := context.WithCancel(l.ctx)
	g, _ := errgroup.WithContext(gCtx)
	g.Go(func() error {
		defer func() {
			if e := recover(); e != nil {
				logx.Error(e)
			}
		}()
		//登陆时查询的用户已经过滤掉软删除用户
		userInfoParam := &systemservice.UserID{ID: res.UserId}
		user, err = l.svcCtx.SystemRpcClient.UserInfo(l.ctx, userInfoParam)
		if err != nil {
			l.Error(err)
			msgErrList.WithMeta("SystemRpcClient.UserInfo", err.Error(), userInfoParam)
		}
		return nil
	})
	g.Go(func() error {
		defer func() {
			if e := recover(); e != nil {
				logx.Error(e)
			}
		}()
		userPageSetParam := &systemservice.UserID{ID: res.UserId}
		userpageset, err = l.svcCtx.SystemRpcClient.UserPageSetInfo(l.ctx, userPageSetParam)
		if err != nil {
			s, _ := status.FromError(err)
			if s.Message() == sql.ErrNoRows.Error() {
				userpageset = new(systemservice.UserPageSet)
				userpageset.ID = 0
				userpageset.UserId = res.UserId
				userpageset.Avatar = ""
				userpageset.DefaultRouter = "dashboard"
				userpageset.SideMode = "#191a23"
				userpageset.ActiveTextColor = "#1890ff"
				userpageset.TextColor = "#fff"
			} else {
				l.Error(err)
				msgErrList.WithMeta("SystemRpcClient.UserPageSetInfo", err.Error(), userPageSetParam)
			}
		}
		return nil
	})
	g.Go(func() error {
		defer func() {
			if e := recover(); e != nil {
				logx.Error(e)
			}
		}()
		getUserRoleByIDParam := &systemservice.UserID{ID: res.UserId}
		userrole, err := l.svcCtx.SystemRpcClient.GetUserRoleByUserID(l.ctx, getUserRoleByIDParam)
		if err != nil {
			l.Error(err)
			msgErrList.WithMeta("SystemRpcClient.GetUserRoleByUserID", err.Error(), getUserRoleByIDParam)
			return nil
		}

		mr.MapReduce(
			func(source chan<- interface{}) {
				for _, v := range userrole.UserRole {
					source <- v.RoleID
				}
			},
			func(item interface{}, writer mr.Writer, cancel func(error)) {
				roleid := item.(uint64)
				roleInfoParam := &systemservice.RoleID{ID: roleid}
				if role, err := l.svcCtx.SystemRpcClient.RoleInfo(l.ctx, roleInfoParam); err != nil {
					l.Error(err)
					msgErrList.WithMeta("SystemRpcClient.RoleInfo", err.Error(), roleInfoParam)
				} else {
					writer.Write(role)
				}
			},
			func(pipe <-chan interface{}, writer mr.Writer, cancel func(error)) {
				for v := range pipe {
					role, ok := v.(*systemservice.Role)
					if ok {
						if role.DeleteTime == 0 {
							rolelist = append(rolelist, role.Role)
						}
					} else {
						logx.Errorf("mr reducer断言失败, 实际类型: (%T), 断言类型: (*systemservice.Role)", v, v)
					}
				}
			},
		)
		return nil
	})
	g.Wait()

	var msg string = "OK"
	if len(msgErrList.List) != 0 {
		msg = fmt.Sprintf("Not OK(%d)", len(msgErrList.List))
	}
	return &types.LoginResponse{
		HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: msg, Meta: msgErrList.List},
		Token:              res.GetToken(),
		ExpireAt:           res.GetExporeAt(),
		RefreshAfter:       res.GetRefreshAfter(),
		Name:               user.Name,
		UserPageSet: types.UserPageSet{
			Name:            user.Name,
			NickName:        user.NickName,
			Avatar:          userpageset.Avatar,
			DefaultRouter:   userpageset.DefaultRouter,
			SideMode:        userpageset.SideMode,
			ActiveTextColor: userpageset.ActiveTextColor,
			TextColor:       userpageset.TextColor,
		},
		Role: rolelist,
	}, nil
}
