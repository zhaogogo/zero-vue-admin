package role

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AllRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAllRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AllRoleLogic {
	return &AllRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AllRoleLogic) AllRole() (resp *types.AllRoleResponse, err error) {
	proles, err := l.svcCtx.SystemRpcClient.AllRoleList(l.ctx, &systemservice.Empty{})
	if err != nil {
		return nil, err
	}
	roles := []types.Role{}
	for _, prole := range proles.List {
		role := types.Role{
			ID:   prole.ID,
			Role: prole.Role,
			Name: prole.Name,
		}
		roles = append(roles, role)
	}

	return &types.AllRoleResponse{
		HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: "OK"},
		Total:              len(roles),
		List:               roles,
	}, nil
}
