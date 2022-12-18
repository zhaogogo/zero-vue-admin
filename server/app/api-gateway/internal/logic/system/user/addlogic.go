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

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.AddUserRequest) (resp *types.HttpCommonResponse, err error) {
	username := l.ctx.Value("userName").(string)
	param := &systemservice.AddUserAndUserRoleRequest{
		User: &systemservice.User{
			ID:         0,
			Name:       req.Name,
			NickName:   req.NickName,
			PassWord:   req.PassWord,
			UserType:   req.UserType,
			Email:      req.Email,
			Phone:      strconv.FormatInt(req.Phone, 10),
			Department: req.Department,
			Position:   req.Position,
			CreateBy:   username,
			UpdateBy:   username,
			DeleteBy:   "",
		},
		RoleList: req.RoleList,
	}
	_, err = l.svcCtx.SystemRpcClient.AddUserAndUserRole(l.ctx, param)
	if err != nil {
		return nil, errorx.NewByCode(err, errorx.GRPC_ERROR).WithMeta("*SystemRpcClient.AddUserAndUserRole", err.Error(), param)
	}

	return &types.HttpCommonResponse{Code: 200, Msg: "OK"}, nil
}