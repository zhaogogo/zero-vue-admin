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

type APIPagingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAPIPagingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *APIPagingLogic {
	return &APIPagingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *APIPagingLogic) APIPaging(in *pb.APIPagingRequest) (*pb.APIPagingResponse, error) {
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
	apis, err := l.svcCtx.APIModel.FindPaging_NC(l.ctx, param)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, errors.Wrap(err, "数据库查询错误")
	}
	papis := []*pb.API{}
	for _, api := range apis {
		papi := &pb.API{
			ID:       api.Id,
			API:      api.Api,
			Group:    api.Group,
			Describe: api.Describe,
			Method:   api.Method,
		}
		papis = append(papis, papi)
	}

	return &pb.APIPagingResponse{
		APIs: papis,
	}, nil
}
