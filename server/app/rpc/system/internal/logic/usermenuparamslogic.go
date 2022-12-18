package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserMenuParamsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserMenuParamsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserMenuParamsLogic {
	return &UserMenuParamsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserMenuParamsLogic) UserMenuParams(in *pb.UserID) (*pb.UserMenuParamsResponse, error) {
	userMenuParams, err := l.svcCtx.UserMenuParamsModel.FindByUserID(l.ctx, l.svcCtx.Redis, in.ID)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return nil, err
		}
		return nil, errors.Wrap(err, "数据库错误")
	}
	pbUserMenuParams := []*pb.UserMenuParams{}
	for _, p := range userMenuParams {
		pbUserMenuParams = append(pbUserMenuParams, &pb.UserMenuParams{
			ID:     p.Id,
			UserID: p.UserId,
			MenuID: p.MenuId,
			Type:   p.Type,
			Key:    p.Key,
			Value:  p.Value,
		})
	}
	return &pb.UserMenuParamsResponse{
		UserMenuParams: pbUserMenuParams,
	}, nil
}
