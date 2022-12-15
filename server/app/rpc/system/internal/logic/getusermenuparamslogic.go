package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserMenuParamsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserMenuParamsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserMenuParamsLogic {
	return &GetUserMenuParamsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserMenuParamsLogic) GetUserMenuParams(in *pb.UserID) (*pb.UserMenuParamsList, error) {
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
	return &pb.UserMenuParamsList{
		UserMenuParams: pbUserMenuParams,
	}, nil
}
