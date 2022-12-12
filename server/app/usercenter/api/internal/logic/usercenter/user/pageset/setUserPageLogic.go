package pageset

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	errorx2 "github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/usercenter/api/internal/pkg/logiccommon"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/usercenter/api/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/usercenter/api/internal/types"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/usercenter/model"
)

type SetUserPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetUserPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUserPageLogic {
	return &SetUserPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetUserPageLogic) SetUserPage(req *types.UserPageRequest) (resp *types.UserPageSetResponse, err error) {
	userid, err := logiccommon.GetUserIdWithJWT(l.ctx)
	if err != nil {
		return nil, errorx2.New(errorx2.SERVER_COMMON_ERROR, "获取用户ID失败")
	}
	logx.Field("userID", userid)

	err = l.svcCtx.UserPageSetModel.CreateODuplicateByUserId(
		l.ctx,
		&model.UserPageSet{UserId: userid, Avatar: req.Avatar, DefaultRouter: req.DefaultRouter, SideMode: req.SideMode, TextColor: req.TextColor, ActiveTextColor: req.ActiveTextColor},
		userid,
	)
	if err != nil {
		logx.Errorf("数据库错误, error: %v", err)
		return nil, errorx2.NewFromCode(errorx2.DB_ERROR, err)
	}
	userpageset, err := l.svcCtx.UserPageSetModel.FindOneByUserId(l.ctx, userid)
	if err != nil {
		return nil, errorx2.NewFromCode(errorx2.DB_ERROR, err)
	}
	return &types.UserPageSetResponse{
		HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: "OK"},
		Avatar:             userpageset.Avatar,
		DefaultRouter:      userpageset.DefaultRouter,
		SideMode:           userpageset.SideMode,
		ActiveTextColor:    userpageset.ActiveTextColor,
		TextColor:          userpageset.TextColor,
	}, nil
}
