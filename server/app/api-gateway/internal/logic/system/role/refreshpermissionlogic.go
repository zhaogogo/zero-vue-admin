package role

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshPermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshPermissionLogic {
	return &RefreshPermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshPermissionLogic) RefreshPermission() (resp *types.HttpCommonResponse, err error) {
	_, err = l.svcCtx.SystemRpcClient.RefreshCasbinPolicy(l.ctx, &systemservice.Empty{})
	if err != nil {
		return nil, errorx.New(err, "casbin刷新Policy错误").WithMeta("SystemRpcClient.RefreshCasbinPolicy", err.Error(), "")
	}

	return &types.HttpCommonResponse{Code: 200, Msg: "OK"}, nil
}
