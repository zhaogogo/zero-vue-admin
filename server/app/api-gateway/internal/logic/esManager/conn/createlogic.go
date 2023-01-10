package conn

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/esManager/esmanagerservice"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.ESConnCreateRequest) (resp *types.HttpCommonResponse, err error) {
	createEsConnParam := &esmanagerservice.CreateESConnRequest{
		ESConn:   req.ESConn,
		Version:  req.Version,
		User:     req.User,
		PassWord: req.PassWord,
		Describe: req.Describe,
	}
	_, err = l.svcCtx.ESManagerRpcClient.CreateESConn(l.ctx, createEsConnParam)
	if err != nil {
		return nil, errorx.New(err, "创建ES连接失败").WithMeta("ESManagerRpcClient.CreateESConn", err.Error(), createEsConnParam)
	}

	return &types.HttpCommonResponse{Code: 200, Msg: "OK"}, nil
}
