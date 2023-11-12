package hosts

import (
	"context"
	"fmt"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/responseerror/errorx"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSlienceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSlienceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSlienceLogic {
	return &GetSlienceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSlienceLogic) GetSlience(req *types.SlienceGetRequest) (resp *types.SlienceGetResponse, err error) {
	host := types.Host{}
	err = l.svcCtx.MonitoringDB().Where(types.Host{Host: req.Host}).Take(&host).Error
	if err != nil {
		return nil, errorx.New(err, fmt.Sprintf("获取host %s id 失败", req.Host))
	}
	alertSlience := []types.SlienceName{}
	err = l.svcCtx.MonitoringDB().Where(types.SlienceName{HostID: host.Id}).Preload("Matchers", "host_id = ?", host.Id).Find(&alertSlience).Error
	if err != nil {
		return nil, errorx.New(err, fmt.Sprintf("获取host %s 静默规则失败", req.Host))
	}
	return &types.SlienceGetResponse{
		HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: "ok"},
		HostSliences: types.HostSliences{
			ID:       host.Id,
			Host:     host.Host,
			Sliences: alertSlience,
		},
	}, nil
}
