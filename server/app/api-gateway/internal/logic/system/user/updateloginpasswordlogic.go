package user

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLoginPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLoginPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLoginPasswordLogic {
	return &UpdateLoginPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLoginPasswordLogic) UpdateLoginPassword(req *types.UpdateLoginPasswordRequest) (resp *types.HttpCommonResponse, err error) {
	userid := l.ctx.Value("user_id").(uint64)
	updateUserPassowrdParam := &systemservice.UpdateUserPasswordRequest{ID: userid, Password: req.Password}
	_, err = l.svcCtx.SystemRpcClient.UpdateUserPassword(l.ctx, updateUserPassowrdParam)
	if err != nil {
		return nil, errorx.NewByCode(err, errorx.GRPC_ERROR).WithMeta("SystemRpcClient.ChangePassword", err.Error(), updateUserPassowrdParam)
	}
	return &types.HttpCommonResponse{
		Code: 200,
		Msg:  "OK",
	}, nil
}
