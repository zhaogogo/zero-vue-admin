package user

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req *types.AddUserRequest) (resp *types.HttpCommonResponse, err error) {
	username := l.ctx.Value("c-UserName").(string)
	_, err = l.svcCtx.SystemRpcClient.AddUserAndUserRole(l.ctx, &systemservice.AddUserAndUserRoleRequest{
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
	})
	if err != nil {
		return nil, err
	}

	return &types.HttpCommonResponse{Code: 200, Msg: "OK"}, nil
}
