package conn

import (
	"context"
	"database/sql"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/esManager/esmanagerservice"
	"google.golang.org/grpc/status"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AllLogic {
	return &AllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AllLogic) All() (resp *types.ESConnResponse, err error) {
	esConnAllParam := &esmanagerservice.Empty{}
	esconns, err := l.svcCtx.ESManagerRpcClient.ESConnAll(l.ctx, esConnAllParam)
	if err != nil {
		s, _ := status.FromError(err)
		if s.Message() == sql.ErrNoRows.Error() {
			return &types.ESConnResponse{
				HttpCommonResponse:   types.HttpCommonResponse{Code: 200, Msg: "OK"},
				PagingCommonResponse: types.PagingCommonResponse{Total: 0},
				List:                 []types.Conn{},
			}, nil
		}
		return nil, errorx.New(err, "获取ES全部连接失败").WithMeta("ESManagerRpcClient.ESConnAll", err.Error(), esConnAllParam)
	}
	v := []types.Conn{}
	for _, esconn := range esconns.ESConns {
		tp := types.Conn{
			ID:       esconn.ID,
			ESConn:   esconn.ESConn,
			Version:  esconn.Version,
			User:     esconn.User,
			PassWord: esconn.PassWord,
			Describe: esconn.Describe,
		}
		v = append(v, tp)
	}

	return &types.ESConnResponse{
		HttpCommonResponse:   types.HttpCommonResponse{Code: 200, Msg: "OK"},
		PagingCommonResponse: types.PagingCommonResponse{Total: 0},
		List:                 v,
	}, nil
}
