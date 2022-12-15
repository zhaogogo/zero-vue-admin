package user

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/utils"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SoftDeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSoftDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SoftDeleteUserLogic {
	return &SoftDeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SoftDeleteUserLogic) SoftDeleteUser(req *types.SoftDeleteUserRequest) (resp *types.HttpCommonResponse, err error) {
	userid, err := utils.GetUserIdWithJWT(l.ctx)
	if err != nil {
		return nil, errorx.NewByCode(err, errorx.PARSE_JWTTOKE_NERROR)
	}
	user, err := l.svcCtx.SystemRpcClient.UserInfo(l.ctx, &systemservice.UserID{ID: userid})
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.SystemRpcClient.SoftDeleteUser(l.ctx, &systemservice.SoftDeleteUserRequest{UserID: req.UserID, DeleteBy: user.Name, State: req.State})
	if err != nil {
		return nil, err
	}
	return &types.HttpCommonResponse{Code: 200, Msg: "OK"}, nil
}
