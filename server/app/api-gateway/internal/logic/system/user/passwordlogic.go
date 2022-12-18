package user

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PasswordLogic {
	return &PasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PasswordLogic) Password(req *types.ChangePasswordRequest) (resp *types.HttpCommonResponse, err error) {
	param := &systemservice.ChangePasswordRequest{ID: req.ID, Password: req.Password}
	_, err = l.svcCtx.SystemRpcClient.ChangePassword(l.ctx, param)
	if err != nil {
		return nil, errorx.NewByCode(err, errorx.GRPC_ERROR).WithMeta("SystemRpcClient.ChangePassword", err.Error(), param)
	}
	return &types.HttpCommonResponse{
		Code: 200,
		Msg:  "OK",
	}, nil
}
