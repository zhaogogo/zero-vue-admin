package conn

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/esManager/esmanagerservice"

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

func (l *DetailLogic) Detail(req *types.ESConnDetailRequest) (resp *types.ESConnDetailResponse, err error) {
	esDetailParam := &esmanagerservice.ESConnID{ID: req.ID}
	pesconn, err := l.svcCtx.ESManagerRpcClient.ESConnDetail(l.ctx, esDetailParam)
	if err != nil {
		return nil, errorx.New(err, "获取ES连接详情失败").WithMeta("ESManagerRpcClient.ESConnDetail", err.Error(), esDetailParam)
	}

	return &types.ESConnDetailResponse{
		HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: "OK"},
		Detail: types.Conn{
			ID:       pesconn.ID,
			ESConn:   pesconn.ESConn,
			Version:  pesconn.Version,
			User:     pesconn.User,
			PassWord: "",
			Describe: pesconn.Describe,
		},
	}, nil
}
