package logic

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/model/system"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type APITotalLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAPITotalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *APITotalLogic {
	return &APITotalLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *APITotalLogic) APITotal(in *pb.APIPagingRequest) (*pb.Total, error) {
	param := &system.PagingAPIList{
		Page:     in.Page,
		PageSize: in.PageSize,
		OrderKey: in.OrderKey,
		Order:    in.Order,
		Api:      in.Api,
		Describe: in.Describe,
		Method:   in.Method,
		Group:    in.Group,
	}
	total, err := l.svcCtx.APIModel.Total_NC(l.ctx, param)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, errors.Wrap(err, "数据库查询失败")
	}

	return &pb.Total{Total: total}, nil
}
