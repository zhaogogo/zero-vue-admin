package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/model/system"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserPageSetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserPageSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserPageSetLogic {
	return &UpdateUserPageSetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserPageSetLogic) UpdateUserPageSet(in *pb.UpdateUserPageSetRequest) (*pb.Empty, error) {
	pagesetInfo, err := l.svcCtx.UserPageSetModel.FindOneByUserId(l.ctx, in.UserID)
	if err != nil {
		if err == sqlc.ErrNotFound {
			pageset := &system.UserPageSet{
				UserId:          in.UserID,
				Avatar:          in.Avatar,
				DefaultRouter:   in.DefaultRouter,
				SideMode:        in.SideMode,
				TextColor:       in.TextColor,
				ActiveTextColor: in.ActiveTextColor,
			}
			_, err := l.svcCtx.UserPageSetModel.Insert(l.ctx, pageset)
			if err != nil {
				return nil, err
			} else {
				return &pb.Empty{}, nil
			}
		}
		return nil, errors.Wrap(err, "数据库错误")
	}
	err = l.svcCtx.UserPageSetModel.Update(l.ctx, &system.UserPageSet{
		Id:              pagesetInfo.Id,
		UserId:          in.UserID,
		Avatar:          in.Avatar,
		DefaultRouter:   in.DefaultRouter,
		SideMode:        in.SideMode,
		TextColor:       in.TextColor,
		ActiveTextColor: in.ActiveTextColor,
	})
	if err != nil {
		return nil, errors.Wrap(err, "数据库错误")
	}
	return &pb.Empty{}, nil
}
