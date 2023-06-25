package connect

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/model/monitoring"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PagingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPagingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PagingLogic {
	return &PagingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PagingLogic) Paging(req *types.ConnectManagerPagingRequest) (resp *types.ConnectManagerPagingResponse, err error) {
	connectmanager := []monitoring.StoreConnectManager{}
	if err := l.svcCtx.MonitoringDB().Find(&connectmanager).Limit(int(req.PageSize)).Offset(int((req.Page - 1) * req.PageSize)).Error; err != nil {
		return nil, errorx.New(err, "获取连接列表失败")
	}

	var total int64
	if err := l.svcCtx.MonitoringDB().Model(&monitoring.StoreConnectManager{}).Count(&total).Error; err != nil {
		return nil, errorx.New(err, "获取连接列表总数失败")
	}
	res := []types.ConnectManager{}
	for _, v := range connectmanager {
		c := types.ConnectManager{
			ID:        v.ID,
			Type:      v.Type,
			Env:       v.Env,
			Host:      v.Host,
			AccessKey: v.AccessKey,
			SecretKey: v.SecretKey,
		}
		res = append(res, c)
	}
	return &types.ConnectManagerPagingResponse{
		HttpCommonResponse:   types.HttpCommonResponse{Code: 200, Msg: "OK", Meta: nil},
		PagingCommonResponse: types.PagingCommonResponse{Page: req.Page, PageSize: req.PageSize, Total: total},
		List:                 res,
	}, nil
}
