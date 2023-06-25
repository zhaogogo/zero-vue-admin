package connect

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/model/monitoring"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.ConnectManagerUpdateRequest) (resp *types.HttpCommonResponse, err error) {
	newconn := monitoring.StoreConnectManager{ID: req.ID, Type: req.Type, Env: req.Env, Host: req.Host, AccessKey: req.AccessKey, SecretKey: req.SecretKey}
	if err := l.svcCtx.MonitoringDB().Save(newconn).Error; err != nil {
		return nil, errorx.New(err, "更新失败")
	}

	return &types.HttpCommonResponse{Code: 200, Msg: "OK", Meta: nil}, nil
}
