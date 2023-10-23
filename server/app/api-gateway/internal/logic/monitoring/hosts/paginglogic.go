package hosts

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/responseerror/errorx"
	"golang.org/x/sync/errgroup"

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

func (l *PagingLogic) Paging(req *types.HostPagingRequest) (resp *types.HostResponse, err error) {
	var (
		_          errorx.MsgErrList
		hosts      []types.Host
		hostsCount int64
	)
	gCtx, _ := context.WithCancel(l.ctx)
	g, _ := errgroup.WithContext(gCtx)
	g.Go(func() error {
		defer func() {
			if e := recover(); e != nil {
				logx.Error(e)
			}
		}()
		err := l.svcCtx.MonitoringDB().Find(&types.Host{}).Count(&hostsCount).Error
		return err
	})

	g.Go(func() error {
		defer func() {
			if e := recover(); e != nil {
				logx.Error(e)
			}
		}()
		err := l.svcCtx.MonitoringDB().Find(&hosts).Order("id").Limit(int(req.PageSize)).Offset(int((req.Page - 1) * req.PageSize)).Association("Tags").Error
		return err
	})

	err = g.Wait()
	if err != nil {
		return nil, errorx.New(err, "获取hosts失败")
	}
	return &types.HostResponse{
		HttpCommonResponse:   types.HttpCommonResponse{Code: 200, Msg: "OK", Meta: nil},
		PagingCommonResponse: types.PagingCommonResponse{Page: req.Page, PageSize: req.PageSize, Total: hostsCount},
		List:                 hosts,
	}, nil
}
