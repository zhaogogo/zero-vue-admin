package conn

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/esManager/esmanagerservice"
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

func (l *PagingLogic) Paging(req *types.ConnRequest) (resp *types.ConnResponse, err error) {
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
		totalParam := &esmanagerservice.Empty{}
		ptotal, err := l.svcCtx.ESManagerRpcClient.ESConnTotal(l.ctx, totalParam)
		if err != nil {
			msgErrList.WithMeta("ESManagerRpcClient.ESConnTotal", err.Error(), totalParam)
		}
		if ptotal != nil {
			total = ptotal.Total
		}
	}()

	esConnPagingParam := &esmanagerservice.ESConnPagingRequest{Page: req.Page, PageSize: req.PageSize}
	conns, err := l.svcCtx.ESManagerRpcClient.ESConnPaging(l.ctx, esConnPagingParam)
	if err != nil {
		s, _ := status.FromError(err)
		if s.Message() == sql.ErrNoRows.Error() {
			msgErrList.WithMeta("ESManagerRpcClient.ESConnPaging", err.Error(), esConnPagingParam)
			return &types.ConnResponse{
				HttpCommonResponse:   types.HttpCommonResponse{Code: 200, Msg: "OK", Meta: msgErrList.List},
				PagingCommonResponse: types.PagingCommonResponse{Page: req.Page, PageSize: req.PageSize, Total: total},
				List:                 []types.Conn{},
			}, nil
		}
		return nil, errorx.NewByCode(err, errorx.GRPC_ERROR).WithMeta("ESManagerRpcClient.ESConnPaging", err.Error(), esConnPagingParam)
	}
	tconns := []types.Conn{}
	for _, conn := range conns.ESConns {
		t := types.Conn{
			ID:       conn.ID,
			ESConn:   conn.ESConn,
			Version:  conn.Version,
			User:     conn.User,
			PassWord: conn.PassWord,
			Describe: conn.Describe,
		}
		tconns = append(tconns, t)
	}
	wg.Wait()
	var (
		msg     = "OK"
		elcount = len(msgErrList.List)
	)
	if elcount != 0 {
		msg = fmt.Sprintf("Not OK(%d)", elcount)
	}
	return &types.ConnResponse{
		HttpCommonResponse:   types.HttpCommonResponse{Code: 200, Msg: msg, Meta: msgErrList.List},
		PagingCommonResponse: types.PagingCommonResponse{Page: req.Page, PageSize: req.PageSize, Total: total},
		List:                 tconns,
	}, nil
}
