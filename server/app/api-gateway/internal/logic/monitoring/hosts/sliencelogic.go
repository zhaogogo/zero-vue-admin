package hosts

import (
	"context"
	"fmt"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/responseerror/errorx"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SlienceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSlienceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SlienceLogic {
	return &SlienceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SlienceLogic) Slience(req *types.SlienceRequest) (resp *types.SlienceResponse, err error) {
	host := types.Host{}
	err = l.svcCtx.MonitoringDB().Take(&host).Error
	if err != nil {
		return nil, errorx.New(err, fmt.Sprintf("获取host %s id 失败", req.Host))
	}
	alertSlience := []types.SlienceName{}
	err = l.svcCtx.MonitoringDB().Where(types.SlienceName{HostID: host.Id}).Preload("Matchers", "host_id = ?", host.Id).Find(&alertSlience).Error
	if err != nil {
		return nil, errorx.New(err, fmt.Sprintf("获取host %s 静默规则失败", req.Host))
	}
	return &types.SlienceResponse{
		HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: "ok"},
		List:               alertSlience,
	}, nil
}
