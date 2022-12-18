package user

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSoftLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSoftLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSoftLogic {
	return &DeleteSoftLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSoftLogic) DeleteSoft(req *types.SoftDeleteUserRequest) (resp *types.HttpCommonResponse, err error) {
	username := l.ctx.Value("userName").(string)
	param := &systemservice.SoftDeleteUserRequest{UserID: req.UserID, DeleteBy: username, State: req.State}
	_, err = l.svcCtx.SystemRpcClient.SoftDeleteUser(l.ctx, param)
	if err != nil {
		return nil, errorx.NewByCode(err, errorx.GRPC_ERROR).WithMeta("*SystemRpcClient.SoftDeleteUser", err.Error(), param)
	}
	return &types.HttpCommonResponse{Code: 200, Msg: "OK"}, nil
}
