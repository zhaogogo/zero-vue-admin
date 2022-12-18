package user

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"
	"strconv"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditLogic {
	return &EditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditLogic) Edit(req *types.EditUserInfoRequest) (resp *types.HttpCommonResponse, err error) {
	loginUser := l.ctx.Value("userName").(string)
	param := &systemservice.EditUserInfoRequest{
		ID:         req.ID,
		Name:       req.Name,
		NickName:   req.NickName,
		Email:      req.Email,
		Phone:      strconv.FormatInt(req.Phone, 10),
		Department: req.Department,
		Position:   req.Position,
		UpdateBy:   loginUser,
	}
	_, err = l.svcCtx.SystemRpcClient.EditUserInfo(l.ctx, param)
	if err != nil {
		return nil, errorx.NewByCode(err, errorx.GRPC_ERROR).WithMeta("SystemRpcClient.EditUserInfo", err.Error(), param)
	}
	return &types.HttpCommonResponse{Code: 200, Msg: "OK"}, nil
}
