package connect

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/model/monitoring"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.ConnectManagerDetailRequest) (resp *types.ConnectManagerDetailResponse, err error) {
	r := monitoring.StoreConnectManager{ID: req.ID}
	if err = l.svcCtx.MonitoringDB().Find(&r).Error; err != nil {
		return nil, errorx.New(err, "获取详情失败")
	}
	return &types.ConnectManagerDetailResponse{
		HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: "200", Meta: nil},
		Detail: types.ConnectManager{
			ID:        r.ID,
			Type:      r.Type,
			Env:       r.Env,
			Host:      r.Host,
			AccessKey: r.AccessKey,
			SecretKey: r.SecretKey,
		},
	}, nil
}
