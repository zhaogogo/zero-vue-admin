package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/model"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAPIMultipleAndCasbinLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteAPIMultipleAndCasbinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAPIMultipleAndCasbinLogic {
	return &DeleteAPIMultipleAndCasbinLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteAPIMultipleAndCasbinLogic) DeleteAPIMultipleAndCasbin(in *pb.DeleteAPIMultipleAndCasbinRequest) (*pb.Empty, error) {
	apis := []model.APIDeleteMultiple{}
	for _, api := range in.APIs {
		apis = append(apis, model.APIDeleteMultiple{
			ID:     api.ID,
			API:    api.Api,
			Method: api.Method,
		})
	}
	err := l.svcCtx.APIModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		err := l.svcCtx.APIModel.TransDeleteMultiple(ctx, session, apis)
		if err != nil {
			return errors.Wrap(err, "删除API模型失败")
		}
		err = l.svcCtx.CasbinRuleModel.TransDeleteMultiple(ctx, session, apis)
		if err != nil {
			return errors.Wrap(err, "删除casbin_rule模型失败")
		}

		err = l.svcCtx.SyncedEnforcer.LoadPolicy()
		if err != nil {
			return errors.Wrap(err, "SyncedEnforcer加载策略失败")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
