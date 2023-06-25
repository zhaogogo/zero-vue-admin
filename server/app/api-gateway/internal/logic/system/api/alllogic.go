package api

import (
	"context"
	"database/sql"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"
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

func (l *AllLogic) All() (resp *types.APIAllResponse, err error) {
	param := &systemservice.Empty{}
	allapi, err := l.svcCtx.SystemRpcClient.APIAll(l.ctx, param)
	if err != nil {
		s, _ := status.FromError(err)
		if s.Message() == sql.ErrNoRows.Error() {
			return &types.APIAllResponse{
				HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: "OK", Meta: err.Error()},
				Total:              0,
				List:               []types.API{},
			}, nil
		}
		return nil, errorx.New(err, "获取全量API失败").WithMeta("SystemRpcClient.APIAll", err.Error(), param)
	}
	tApis := []types.API{}
	for _, papi := range allapi.APIs {
		tApis = append(tApis, types.API{
			ID:       papi.ID,
			API:      papi.API,
			Group:    papi.Group,
			Describe: papi.Describe,
			Method:   papi.Method,
		})
	}
	return &types.APIAllResponse{
		HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: "OK"},
		Total:              int64(len(tApis)),
		List:               tApis,
	}, nil
}
