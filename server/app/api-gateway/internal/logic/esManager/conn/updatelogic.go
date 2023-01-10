package conn

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/esManager/esmanagerservice"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.ESConnUpdateRequest) (resp *types.HttpCommonResponse, err error) {
	updateESConnParam := &esmanagerservice.UpdateESConnRequest{
		ID:       req.ID,
		ESConn:   req.ESConn,
		Version:  req.Version,
		User:     req.User,
		PassWord: req.PassWord,
		Describe: req.Describe,
	}
	_, err = l.svcCtx.ESManagerRpcClient.UpdateESConn(l.ctx, updateESConnParam)
	if err != nil {
		return nil, errorx.New(err, "更新ES连接失败").WithMeta("ESManagerRpcClient.UpdateESConn", err.Error(), updateESConnParam)
	}
	return &types.HttpCommonResponse{Code: 200, Msg: "OK"}, nil
}
