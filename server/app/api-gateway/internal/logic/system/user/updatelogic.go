package user

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"
	"strconv"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UserUpdateRequest) (resp *types.HttpCommonResponse, err error) {
	loginUser := l.ctx.Value("userName").(string)
	param := &systemservice.UpdateUserRequest{
		ID:         req.ID,
		Name:       req.Name,
		NickName:   req.NickName,
		Email:      req.Email,
		Phone:      strconv.FormatInt(req.Phone, 10),
		Department: req.Department,
		Position:   req.Position,
		UpdateBy:   loginUser,
	}
	_, err = l.svcCtx.SystemRpcClient.UpdateUser(l.ctx, param)
	if err != nil {
		return nil, errorx.New(err, "更新用户失败").WithMeta("SystemRpcClient.EditUserInfo", err.Error(), param)
	}
	return &types.HttpCommonResponse{Code: 200, Msg: "OK"}, nil
}
