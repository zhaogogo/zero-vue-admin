package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/model/system"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAPILogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAPILogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAPILogic {
	return &UpdateAPILogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateAPILogic) UpdateAPI(in *pb.API) (*pb.Empty, error) {
	newAPi := &system.Api{
		Id:       in.ID,
		Api:      in.API,
		Group:    in.Group,
		Describe: in.Describe,
		Method:   in.Method,
	}
	oldapi, err := l.svcCtx.APIModel.FindOne(l.ctx, in.ID)
	if err != nil {
		return nil, errors.Wrap(err, "查询API失败")
	}
	err = l.svcCtx.APIModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		err := l.svcCtx.CasbinRuleModel.TransUpdateV2V3(ctx, session, newAPi, oldapi)
		if err != nil {
			return errors.Wrap(err, "更新casbin_rule失败")
		}
		err = l.svcCtx.APIModel.TransUpdate(l.ctx, session, newAPi)
		if err != nil {
			return errors.Wrap(err, "更新API失败")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	l.svcCtx.SyncedEnforcer.LoadPolicy()
	return &pb.Empty{}, nil
}
