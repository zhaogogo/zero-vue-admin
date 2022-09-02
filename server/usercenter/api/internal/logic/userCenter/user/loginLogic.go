package user

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/api/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/api/internal/types"
	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	user, err := l.svcCtx.SystemUserRpc.Login(l.ctx, &pb.LoginRequest{
		UserName: req.UserName,
		PassWord: req.PassWord,
	})
	if err != nil {
		return nil, err
	}
	//保存登陆日志

	return &types.LoginResponse{
		ID:           user.ID,
		UserName:     user.UserName,
		Token:        user.Token,
		ExpireAt:     user.ExpireAt,
		RefreshAfter: user.RefreshAfter,
	}, nil
}
