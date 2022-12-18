package role

import (
	"context"
	"database/sql"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"
	"google.golang.org/grpc/status"

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
	roleList := []types.Role{}
	param := &systemservice.Empty{}
	proles, err := l.svcCtx.SystemRpcClient.AllRoleList(l.ctx, param)
	if err != nil {
		s, _ := status.FromError(err)
		if s.Message() == sql.ErrNoRows.Error() {
			return &types.AllRoleResponse{
				HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: "OK"},
				Total:              0,
				List:               roleList,
			}, nil
		}
		return nil, errorx.NewByCode(err, errorx.GRPC_ERROR).WithMeta("*SystemRpcClient.AllRoleList", err.Error(), param)
	}

	for _, prole := range proles.List {
		role := types.Role{
			ID:   prole.ID,
			Role: prole.Role,
			Name: prole.Name,
		}
		roleList = append(roleList, role)
	}

	return &types.AllRoleResponse{
		HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: "OK"},
		Total:              len(roleList),
		List:               roleList,
	}, nil
}
