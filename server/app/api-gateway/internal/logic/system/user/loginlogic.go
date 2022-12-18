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
		userpageset *systemservice.UserPageSetResponse
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
		userDetailParam := &systemservice.UserID{ID: res.UserId}
		user, err = l.svcCtx.SystemRpcClient.UserDetail(l.ctx, userDetailParam)
		if err != nil {
			l.Error(err)
			msgErrList.WithMeta("SystemRpcClient.UserInfo", err.Error(), userDetailParam)
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
		userpageset, err = l.svcCtx.SystemRpcClient.UserPageSet(l.ctx, userPageSetParam)
		if err != nil {
			s, _ := status.FromError(err)
			if s.Message() == sql.ErrNoRows.Error() {
				userpageset = new(systemservice.UserPageSetResponse)
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
		userRoleByUserIDParam := &systemservice.UserID{ID: res.UserId}
		userrole, err := l.svcCtx.SystemRpcClient.UserRoleByUserID(l.ctx, userRoleByUserIDParam)
		if err != nil {
			l.Error(err)
			msgErrList.WithMeta("SystemRpcClient.GetUserRoleByUserID", err.Error(), userRoleByUserIDParam)
			return nil
		}

		mr.MapReduce(
			func(source chan<- interface{}) {
				for _, v := range userrole.UserRoles {
					source <- v.RoleID
				}
			},
			func(item interface{}, writer mr.Writer, cancel func(error)) {
				roleid := item.(uint64)
				roleDetailParam := &systemservice.RoleID{ID: roleid}
				if role, err := l.svcCtx.SystemRpcClient.RoleDetail(l.ctx, roleDetailParam); err != nil {
					l.Error(err)
					msgErrList.WithMeta("SystemRpcClient.RoleInfo", err.Error(), roleDetailParam)
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

/*
2022-12-19T03:19:55.128+08:00    stat   p2c - conn: 127.0.0.1:8080, load: 15939, reqs: 1        caller=p2c/p2c.go:181
=====> 1 /api/v1/system/user/soft/4 DELETE true
2022-12-19T03:19:55.242+08:00    info   [HTTP]  200  -  DELETE  /api/v1/system/user/soft/4 - 127.0.0.1 - Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36       duration=372.7ms        trace=5ed0dea9cbfdb6724f5068d73ad8a6d1  span=5cda9c7fcdc9808f   caller=handler/loghandler.go:197
2022-12-19T03:19:55.415+08:00    info   [HTTP]  401  -  POST  /api/v1/system/user/paging - 127.0.0.1 - Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 duration=30.5ms trace=cfa2d984dfd0b31268c195f0988b8ecf  span=a46095f18a1ff3bd   caller=handler/loghandler.go:197
==================
WARNING: DATA RACE
Write at 0x00c000491a60 by goroutine 83:
  github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/logic/system/user.(*LoginLogic).Login.func2()
      /Users/zhaoqiang/Documents/owner项目/zero-vue-admin/server/app/api-gateway/internal/logic/system/user/loginlogic.go:69 +0x204
  golang.org/x/sync/errgroup.(*Group).Go.func1()
      /Users/zhaoqiang/go/pkg/mod/golang.org/x/sync@v0.0.0-20210220032951-036812b2e83c/errgroup/errgroup.go:57 +0x68

Previous write at 0x00c000491a60 by goroutine 82:
  github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/logic/system/user.(*LoginLogic).Login.func1()
      /Users/zhaoqiang/Documents/owner项目/zero-vue-admin/server/app/api-gateway/internal/logic/system/user/loginlogic.go:55 +0x200
  golang.org/x/sync/errgroup.(*Group).Go.func1()
      /Users/zhaoqiang/go/pkg/mod/golang.org/x/sync@v0.0.0-20210220032951-036812b2e83c/errgroup/errgroup.go:57 +0x68

Goroutine 83 (running) created at:
  golang.org/x/sync/errgroup.(*Group).Go()
      /Users/zhaoqiang/go/pkg/mod/golang.org/x/sync@v0.0.0-20210220032951-036812b2e83c/errgroup/errgroup.go:54 +0x68
  github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/logic/system/user.(*LoginLogic).Login()
      /Users/zhaoqiang/Documents/owner项目/zero-vue-admin/server/app/api-gateway/internal/logic/system/user/loginlogic.go:62 +0xa38
  github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/handler/system/user.LoginHandler.func1()
      /Users/zhaoqiang/Documents/owner项目/zero-vue-admin/server/app/api-gateway/internal/handler/system/user/loginhandler.go:25 +0x3ac
  net/http.HandlerFunc.ServeHTTP()
      /usr/local/go/src/net/http/server.go:2046 +0x48
  github.com/zeromicro/go-zero/rest/handler.GunzipHandler.func1()
      /Users/zhaoqiang/go/pkg/mod/github.com/zeromicro/go-zero@v1.4.0/rest/handler/gunziphandler.go:26 +0x16c
  net/http.HandlerFunc.ServeHTTP()
      /usr/local/go/src/net/http/server.go:2046 +0x48
  github.com/zeromicro/go-zero/rest/handler.MaxBytesHandler.func2.1()
      /Users/zhaoqiang/go/pkg/mod/github.com/zeromicro/go-zero@v1.4.0/rest/handler/maxbyteshandler.go:24 +0x154
  net/http.HandlerFunc.ServeHTTP()
      /usr/local/go/src/net/http/server.go:2046 +0x48
  github.com/zeromicro/go-zero/rest/handler.MetricHandler.func1.1()
      /Users/zhaoqiang/go/pkg/mod/github.com/zeromicro/go-zero@v1.4.0/rest/handler/metrichandler.go:21 +0xe0
  net/http.HandlerFunc.ServeHTTP()
      /usr/local/go/src/net/http/server.go:2046 +0x48
  github.com/zeromicro/go-zero/rest/handler.RecoverHandler.func1()
      /Users/zhaoqiang/go/pkg/mod/github.com/zeromicro/go-zero@v1.4.0/rest/handler/recoverhandler.go:21 +0xa0
  net/http.HandlerFunc.ServeHTTP()
      /usr/local/go/src/net/http/server.go:2046 +0x48
  github.com/zeromicro/go-zero/rest/handler.(*timeoutHandler).ServeHTTP.func1()
      /Users/zhaoqiang/go/pkg/mod/github.com/zeromicro/go-zero@v1.4.0/rest/handler/timeouthandler.go:79 +0x98

Goroutine 82 (f  golang.org/x/sync/errgroup.(*Group).Go()
      /Users/zhaoqiang/go/pkg/mod/golang.org/x/sync@v0.0.0-20210220032951-036812b2e83c/errgroup/errgroup.go:54 +0x68
  github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/logic/system/user.(*LoginLogic).Login()
      /Users/zhaoqiang/Documents/owner项目/zero-vue-admin/server/app/api-gateway/internal/logic/system/user/loginlogic.go:47 +0x8b8
  github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/handler/system/user.LoginHandler.func1()
      /Users/zhaoqiang/Documents/owner项目/zero-vue-admin/server/app/api-gateway/internal/handler/system/user/loginhandler.go:25 +0x3ac
  net/http.HandlerFunc.ServeHTTP()
      /usr/local/go/src/net/http/server.go:2046 +0x48
  github.com/zeromicro/go-zero/rest/handler.GunzipHandler.func1()
      /Users/zhaoqiang/go/pkg/mod/github.com/zeromicro/go-zero@v1.4.0/rest/handler/gunziphandler.go:26 +0x16c
  net/http.HandlerFunc.ServeHTTP()
      /usr/local/go/src/net/http/server.go:2046 +0x48
  github.com/zeromicro/go-zero/rest/handler.MaxBytesHandler.func2.1()
      /Users/zhaoqiang/go/pkg/mod/github.com/zeromicro/go-zero@v1.4.0/rest/handler/maxbyteshandler.go:24 +0x154
  net/http.HandlerFunc.ServeHTTP()
      /usr/local/go/src/net/http/server.go:2046 +0x48
  github.com/zeromicro/go-zero/rest/handler.MetricHandler.func1.1()
      /Users/zhaoqiang/go/pkg/mod/github.com/zeromicro/go-zero@v1.4.0/rest/handler/metrichandler.go:21 +0xe0
  net/http.HandlerFunc.ServeHTTP()
      /usr/local/go/src/net/http/server.go:2046 +0x48
  github.com/zeromicro/go-zero/rest/handler.RecoverHandler.func1()
      /Users/zhaoqiang/go/pkg/mod/github.com/zeromicro/go-zero@v1.4.0/rest/handler/recoverhandler.go:21 +0xa0
  net/http.HandlerFunc.ServeHTTP()
      /usr/local/go/src/net/http/server.go:2046 +0x48
  github.com/zeromicro/go-zero/rest/handler.(*timeoutHandler).ServeHTTP.func1()
      /Users/zhaoqiang/go/pkg/mod/github.com/zeromicro/go-zero@v1.4.0/rest/handler/timeouthandler.go:79 +0x98
==================

*/
