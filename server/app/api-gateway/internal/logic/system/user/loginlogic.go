package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"
	"golang.org/x/sync/errgroup"
	"strings"
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
	res, err := l.svcCtx.SystemRpcClient.Login(l.ctx, &systemservice.LoginRequest{Name: req.UserName, PassWord: req.PassWord})
	if err != nil {
		return nil, err
	}

	var (
		userpageset *systemservice.UserPageSet
		user        *systemservice.User
		rolelist           = []string{}
		msgErrList         = errorx.MsgErrList{}
		msg         string = "OK"
	)
	gCtx, _ := context.WithCancel(l.ctx)
	g, _ := errgroup.WithContext(gCtx)
	g.Go(func() error {
		//登陆时查询的用户已经过滤掉软删除用户
		user, err = l.svcCtx.SystemRpcClient.UserInfo(l.ctx, &systemservice.UserID{ID: res.UserId})
		if err != nil {
			l.Error(err)
			msgErrList.Append(err.Error())
		}
		return nil
	})
	g.Go(func() error {
		userpageset, err = l.svcCtx.SystemRpcClient.UserPageSetInfo(l.ctx, &systemservice.UserID{ID: res.UserId})
		if err != nil {
			userpageset.ID = 0
			userpageset.UserId = res.UserId
			userpageset.Avatar = ""
			userpageset.DefaultRouter = "dashboard"
			userpageset.SideMode = "#191a23"
			userpageset.ActiveTextColor = "#1890ff"
			userpageset.TextColor = "#fff"

			l.Error(err)
			msgErrList.Append(err.Error())
		}
		return nil
	})
	g.Go(func() error {
		userrole, err := l.svcCtx.SystemRpcClient.GetUserRoleByUserID(l.ctx, &systemservice.UserID{ID: res.UserId})
		if err != nil {
			l.Error(err)
			msgErrList.Append(err.Error())
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
				if role, err := l.svcCtx.SystemRpcClient.RoleInfo(l.ctx, &systemservice.RoleID{ID: roleid}); err != nil {
					l.Error(err)
					msgErrList.Append(err.Error())
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
	if len(msgErrList.List) != 0 {
		msg = strings.Join(msgErrList.List, " | ")
	}
	return &types.LoginResponse{
		HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: msg},
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
