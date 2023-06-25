package conn

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/esManager/esmanagerservice"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DELETELogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDELETELogic(ctx context.Context, svcCtx *svc.ServiceContext) *DELETELogic {
	return &DELETELogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DELETELogic) DELETE(req *types.ESConnDeleteRequest) (resp *types.HttpCommonResponse, err error) {
	deleteEsConnParam := &esmanagerservice.ESConnID{ID: req.ID}
	_, err = l.svcCtx.ESManagerRpcClient.DeleteESConn(l.ctx, deleteEsConnParam)
	if err != nil {
		return nil, errorx.New(err, "删除ES连接失败").WithMeta("ESManagerRpcClient.DeleteESConn", err.Error(), deleteEsConnParam)
	}
	return &types.HttpCommonResponse{Code: 200, Msg: "OK"}, nil
}
