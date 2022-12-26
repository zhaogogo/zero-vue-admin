package logic

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type APIDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAPIDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *APIDetailLogic {
	return &APIDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *APIDetailLogic) APIDetail(in *pb.ApiID) (*pb.API, error) {
	api, err := l.svcCtx.APIModel.FindOne(l.ctx, in.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, errors.Wrap(err, "数据库查询失败")
	}
	return &pb.API{
		ID:       api.Id,
		API:      api.Api,
		Group:    api.Group,
		Describe: api.Describe,
		Method:   api.Method,
	}, nil
}
