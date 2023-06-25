package conn

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/responseerror/errorx"
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

func (l *PingLogic) Ping(req *types.ESConnPingRequest) (resp *types.ESConnPingResponse, err error) {
	pingParam := &esmanagerservice.PingRequest{EsConnID: req.ID}
	res, err := l.svcCtx.ESManagerRpcClient.Ping(l.ctx, pingParam)
	if err != nil {
		s, _ := status.FromError(err)
		return nil, errorx.New(err, s.Message()).WithMeta("ESManagerRpcClient.Ping", err.Error(), pingParam)
	}
	esPingRes := &elastic.PingResult{}
	err = json.Unmarshal(res.Data.Value, esPingRes)
	if err != nil {
		return nil, errorx.New(err, "ESManagerRpcClient.Ping json Unmarshal失败")
	}
	return &types.ESConnPingResponse{
		HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: "OK"},
		Data:               esPingRes,
	}, nil
}
