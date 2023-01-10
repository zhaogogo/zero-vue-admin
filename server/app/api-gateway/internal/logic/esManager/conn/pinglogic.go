package conn

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/esManager/esmanagerservice"
	"google.golang.org/grpc/status"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingLogic) Ping(req *types.PingRequest) (resp *types.PingResponse, err error) {
	pingParam := &esmanagerservice.PingRequest{EsConnID: req.ID}
	res, err := l.svcCtx.ESManagerRpcClient.Ping(l.ctx, pingParam)
	if err != nil {
		s, _ := status.FromError(err)
		return nil, errorx.New(err, s.Message()).WithMeta("ESManagerRpcClient.Ping", err.Error(), pingParam)
	}
	esPingRes := &elastic.PingResult{}
	json.Unmarshal(res.Data.Value, esPingRes)
	return &types.PingResponse{
		HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: "OK"},
		Data:               esPingRes,
	}, nil
}
