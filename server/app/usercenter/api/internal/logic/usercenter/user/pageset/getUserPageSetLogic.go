package pageset

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	errorx2 "github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/usercenter/api/internal/pkg/logiccommon"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/usercenter/api/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/usercenter/api/internal/types"
)

type GetUserPageSetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserPageSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPageSetLogic {
	return &GetUserPageSetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserPageSetLogic) GetUserPageSet() (resp *types.UserPageSetResponse, err error) {
	userid, err := logiccommon.GetUserIdWithJWT(l.ctx)
	if err != nil {
		return nil, errorx2.New(errorx2.SERVER_COMMON_ERROR, "获取用户ID失败")
	}
	user, _ := l.svcCtx.UserModel.FindOne(l.ctx, userid)
	if user == nil {
		user.Name = "Unknow"
		user.NickName = "Unknow"
	}
	userpageset, err := l.svcCtx.UserPageSetModel.FindOneByUserId(l.ctx, userid)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return &types.UserPageSetResponse{
				HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: "OK"},
				Name:               user.Name,
				NickName:           user.NickName,
				Avatar:             "",
				DefaultRouter:      "dashboard",
				SideMode:           "#191a23",
				TextColor:          "#fff",
				ActiveTextColor:    "#1890ff",
			}, nil
		}
		return nil, errors.Wrap(err, "获取角色id失败")
	}

	return &types.UserPageSetResponse{
		HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: "OK"},
		Name:               user.Name,
		NickName:           user.NickName,
		Avatar:             userpageset.Avatar,
		DefaultRouter:      userpageset.DefaultRouter,
		SideMode:           userpageset.SideMode,
		TextColor:          userpageset.TextColor,
		ActiveTextColor:    userpageset.ActiveTextColor,
	}, nil
}
