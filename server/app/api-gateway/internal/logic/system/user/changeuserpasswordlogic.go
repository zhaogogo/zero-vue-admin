package user

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeUserPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeUserPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeUserPasswordLogic {
	return &ChangeUserPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeUserPasswordLogic) ChangeUserPassword(req *types.ChangePasswordRequest) (resp *types.HttpCommonResponse, err error) {
	_, err = l.svcCtx.SystemRpcClient.ChangePassword(l.ctx, &systemservice.ChangePasswordRequest{ID: req.ID, Password: req.Password})
	if err != nil {
		return nil, err
	}
	return &types.HttpCommonResponse{
		Code: 200,
		Msg:  "OK",
	}, nil
}
