package alertrule

import (
	"context"
	"database/sql"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/monitoring/monitoringmanager"
	"google.golang.org/grpc/status"

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

func (l *PagingLogic) Paging(req *types.AlertRulePagingRequest) (resp *types.AlertRulePagingResponse, err error) {
	param := &monitoringmanager.AlertRulePagingRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		OrderKey: req.OrderKey,
		Order:    req.Order,
	}
	res, err := l.svcCtx.MonitoringRpcConf.AlertRulePaging(l.ctx, param)
	if err != nil {
		s, _ := status.FromError(err)
		if s.Message() == sql.ErrNoRows.Error() {
			return &types.AlertRulePagingResponse{
				HttpCommonResponse:   types.HttpCommonResponse{Code: 200, Msg: "OK"},
				PagingCommonResponse: types.PagingCommonResponse{Page: req.Page, PageSize: req.PageSize, Total: 0},
				List:                 []types.AlertRule{},
			}, nil
		}
		return nil, errorx.New(err, "获取AlertRule失败").WithMeta("MonitoringRpcConf.AlertRulePaging", err.Error(), param)
	}
	list := []types.AlertRule{}
	for _, v := range res.AlertRules {
		l := types.AlertRule{
			ID:       v.ID,
			Name:     v.Name,
			Type:     v.Type,
			Group:    v.Group,
			Tag:      v.Tag,
			To:       v.To,
			Expr:     v.Expr,
			Operator: v.Operator,
			Value:    v.Value,
			For:      v.For,
			Summary:  v.Summary,
			Describe: v.Describe,
			IsWrite:  v.IsWrite,
		}
		list = append(list, l)
	}
	return &types.AlertRulePagingResponse{
		HttpCommonResponse:   types.HttpCommonResponse{Code: 200, Msg: "OK"},
		PagingCommonResponse: types.PagingCommonResponse{Page: req.Page, PageSize: req.PageSize, Total: 0},
		List:                 list,
	}, nil
}
