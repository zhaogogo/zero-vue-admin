package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserPageSetInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserPageSetInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserPageSetInfoLogic {
	return &UserPageSetInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserPageSetInfoLogic) UserPageSetInfo(in *pb.UserID) (*pb.UserPageSet, error) {
	pageset, err := l.svcCtx.UserPageSetModel.FindOneByUserId(l.ctx, in.ID)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return nil, err
		}
		return nil, errors.Wrap(err, "数据库错误")
	}

	return &pb.UserPageSet{
		ID:              pageset.Id,
		UserId:          pageset.UserId,
		Avatar:          pageset.Avatar,
		DefaultRouter:   pageset.DefaultRouter,
		SideMode:        pageset.SideMode,
		ActiveTextColor: pageset.ActiveTextColor,
		TextColor:       pageset.TextColor,
	}, nil
	return &pb.UserPageSet{}, nil
}
