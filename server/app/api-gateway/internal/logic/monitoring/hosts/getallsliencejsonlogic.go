package hosts

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/slience"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllSlienceJsonLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllSlienceJsonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllSlienceJsonLogic {
	return &GetAllSlienceJsonLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllSlienceJsonLogic) GetAllSlienceJson() (resp *types.SlienceJsonResponse, err error) {
	err = slience.GetConsumerSliences(l.svcCtx.MonitoringDB(), l.svcCtx.SlienceList)
	if err != nil {
		logx.Errorf("创建Slience，刷新全局规则失败， error: %v", err)
	}
	l.svcCtx.SlienceList.Mu.RLock()
	data := l.svcCtx.SlienceList.Sliences
	l.svcCtx.SlienceList.Mu.RUnlock()

	return &types.SlienceJsonResponse{
		HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: "OK", Meta: nil},
		Data:               data,
	}, nil
}
