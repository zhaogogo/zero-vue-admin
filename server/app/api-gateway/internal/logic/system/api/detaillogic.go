package api

import (
	"context"
	"database/sql"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"
	"google.golang.org/grpc/status"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.APIDetailRequest) (resp *types.APIDetailResponse, err error) {
	apiDetailParam := &systemservice.ApiID{ID: req.ID}
	api, err := l.svcCtx.SystemRpcClient.APIDetail(l.ctx, apiDetailParam)
	if err != nil {
		s, _ := status.FromError(err)
		if s.Message() == sql.ErrNoRows.Error() {
			return nil, errorx.NewByCode(err, errorx.DB_NOTFOUND).WithMeta("SystemRpcClient.APIDetail", err.Error(), apiDetailParam)
		}
		return nil, errorx.NewByCode(err, errorx.GRPC_ERROR).WithMeta("SystemRpcClient.APIDetail", err.Error(), apiDetailParam)
	}
	return &types.APIDetailResponse{
		HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: "OK"},
		Detail: types.API{
			ID:       api.ID,
			API:      api.API,
			Group:    api.Group,
			Describe: api.Describe,
			Method:   api.Method,
		},
	}, nil
}
