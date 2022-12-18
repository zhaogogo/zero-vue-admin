package logic

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AllUserMenuParamsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAllUserMenuParamsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AllUserMenuParamsLogic {
	return &AllUserMenuParamsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AllUserMenuParamsLogic) AllUserMenuParams(in *pb.Empty) (*pb.UserMenuParamsList, error) {
	usermenuparams, err := l.svcCtx.UserMenuParamsModel.FindAll_NC(l.ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, errors.Wrap(err, "数据库错误")
	}
	puserMenuParams := []*pb.UserMenuParams{}
	for _, usermenuparam := range usermenuparams {
		p := &pb.UserMenuParams{
			ID:     usermenuparam.Id,
			UserID: usermenuparam.UserId,
			MenuID: usermenuparam.MenuId,
			Type:   usermenuparam.Type,
			Key:    usermenuparam.Key,
			Value:  usermenuparam.Value,
		}
		puserMenuParams = append(puserMenuParams, p)
	}

	return &pb.UserMenuParamsList{
		UserMenuParams: puserMenuParams,
	}, nil
}
