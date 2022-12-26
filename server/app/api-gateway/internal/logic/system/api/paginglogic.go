package api

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"
	"google.golang.org/grpc/status"
	"sync"

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

func (l *PagingLogic) Paging(req *types.APIPagingRequest) (resp *types.APIPagingResponse, err error) {
	/*
		"order": "ascending"
		"order": "descending"
		"order": ""
	*/
	param := &systemservice.APIPagingRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Order:    req.Order,
		OrderKey: req.OrderKey,
		Api:      req.Api,
		Describe: req.Describe,
		Method:   req.Method,
		Group:    req.Group,
	}
	var total int64
	var msgErrList = errorx.MsgErrList{}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			if r := recover(); r != nil {
				logx.Error(r)
			}
		}()
		ptotal, err := l.svcCtx.SystemRpcClient.APITotal(l.ctx, param)
		if err != nil {
			msgErrList.WithMeta("SystemRpcClient.APITotal", err.Error(), param)
		}
		if ptotal != nil {
			total = ptotal.Total
		}
	}()

	apis, err := l.svcCtx.SystemRpcClient.APIPaging(l.ctx, param)
	if err != nil {
		s, _ := status.FromError(err)
		if s.Message() == sql.ErrNoRows.Error() {
			msgErrList.WithMeta("SystemRpcClient.APIPaging", err.Error(), param)
			return &types.APIPagingResponse{
				HttpCommonResponse:   types.HttpCommonResponse{Code: 200, Msg: "OK", Meta: msgErrList.List},
				PagingCommonResponse: types.PagingCommonResponse{Page: req.Page, PageSize: req.PageSize},
				List:                 []types.API{},
			}, nil
		}
		return nil, errorx.NewByCode(err, errorx.GRPC_ERROR).WithMeta("SystemRpcClient.APIPaging", err.Error(), param)
	}
	tapis := []types.API{}
	for _, api := range apis.APIs {
		t := types.API{
			ID:       api.ID,
			API:      api.API,
			Group:    api.Group,
			Describe: api.Describe,
			Method:   api.Method,
		}
		tapis = append(tapis, t)
	}

	wg.Wait()
	var (
		msg     = "OK"
		elcount = len(msgErrList.List)
	)
	if elcount != 0 {
		msg = fmt.Sprintf("Not OK(%d)", elcount)
	}

	return &types.APIPagingResponse{
		HttpCommonResponse:   types.HttpCommonResponse{Code: 200, Msg: msg, Meta: msgErrList.List},
		PagingCommonResponse: types.PagingCommonResponse{Page: req.Page, PageSize: req.PageSize, Total: total},
		List:                 tapis,
	}, nil
}
