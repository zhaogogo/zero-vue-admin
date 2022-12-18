package user

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.UserDeleteRequest) (resp *types.HttpCommonResponse, err error) {
	params := &systemservice.UserID{ID: req.ID}
	_, err = l.svcCtx.SystemRpcClient.DeleteUser(l.ctx, params)
	if err != nil {
		// "删除用户错误"
		return nil, errorx.NewByCode(err, errorx.GRPC_ERROR).WithMeta("*SystemRpcClient.DeleteUser", err.Error(), params)
	}
	return &types.HttpCommonResponse{Code: 200, Msg: "OK"}, nil
}
