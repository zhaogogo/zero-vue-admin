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

type SetUserPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetUserPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUserPageLogic {
	return &SetUserPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetUserPageLogic) SetUserPage(req *types.UserPageRequest) (resp *types.HttpCommonResponse, err error) {
	//a := l.ctx.Value("c-UserName")
	//fmt.Printf("%T  %v\n", a, a) //string  test1
	userid, err := utils.GetUserIdWithJWT(l.ctx)
	if err != nil {
		return nil, errorx.New(err, errorx.SERVER_COMMON_ERROR, "JWT转换用户ID类型失败")
	}

	_, err = l.svcCtx.SystemRpcClient.SetUserPageSet(l.ctx, &systemservice.SetUserPageSetRequest{
		UserID:          userid,
		Avatar:          req.Avatar,
		DefaultRouter:   req.DefaultRouter,
		SideMode:        req.SideMode,
		ActiveTextColor: req.ActiveTextColor,
		TextColor:       req.TextColor,
	})
	if err != nil {
		return nil, err
	}

	return &types.HttpCommonResponse{
		Code: 200,
		Msg:  "OK",
	}, nil
}
