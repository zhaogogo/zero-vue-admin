package user

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PageLogic {
	return &PageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PageLogic) Page(req *types.UserPageRequest) (resp *types.HttpCommonResponse, err error) {
	userid := l.ctx.Value("user_id").(uint64)
	params := &systemservice.UpdateUserPageSetRequest{
		UserID:          userid,
		Avatar:          req.Avatar,
		DefaultRouter:   req.DefaultRouter,
		SideMode:        req.SideMode,
		ActiveTextColor: req.ActiveTextColor,
		TextColor:       req.TextColor,
	}
	_, err = l.svcCtx.SystemRpcClient.UpdateUserPageSet(l.ctx, params)
	if err != nil {
		return nil, errorx.NewByCode(err, errorx.GRPC_ERROR).WithMeta("*SetUserPageSet", err.Error(), params)
	}

	return &types.HttpCommonResponse{
		Code: 200,
		Msg:  "OK",
	}, nil
	return
}
