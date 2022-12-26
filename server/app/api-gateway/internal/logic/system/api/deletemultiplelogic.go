package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMultipleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMultipleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMultipleLogic {
	return &DeleteMultipleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMultipleLogic) DeleteMultiple(req *types.APIDeleteMultipleRequest) (resp *types.HttpCommonResponse, err error) {
	s, _ := json.MarshalIndent(req, "", "\t")
	fmt.Println(string(s))
	param := &systemservice.DeleteAPIMultipleAndCasbinRequest{}
	for _, api := range req.APIs {
		param.APIs = append(param.APIs, &systemservice.DeleteAPIAndCasbinRequest{
			ID:     api.ID,
			Api:    api.Api,
			Method: api.Method,
		})
	}
	_, err = l.svcCtx.SystemRpcClient.DeleteAPIMultipleAndCasbin(l.ctx, param)
	if err != nil {
		return nil, errorx.NewByCode(err, errorx.GRPC_ERROR).WithMeta("SystemRpcClient.DeleteAPIMultipleAndCasbin", err.Error(), param)
	}
	return &types.HttpCommonResponse{Code: 200, Msg: "OK"}, nil
}
