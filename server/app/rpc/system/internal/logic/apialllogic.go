package logic

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type APIAllLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAPIAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *APIAllLogic {
	return &APIAllLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *APIAllLogic) APIAll(in *pb.Empty) (*pb.APIAllResponse, error) {
	apis, err := l.svcCtx.APIModel.FindAll_NC(l.ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, errors.Wrap(err, "数据库查询错误")
	}
	papis := []*pb.API{}
	for _, api := range apis {
		papis = append(papis, &pb.API{
			ID:       api.Id,
			API:      api.Api,
			Group:    api.Group,
			Describe: api.Describe,
			Method:   api.Method,
		})
	}
	return &pb.APIAllResponse{
		APIs: papis,
	}, nil
}
