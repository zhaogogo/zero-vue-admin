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

type AllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AllLogic {
	return &AllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AllLogic) All() (resp *types.RoleAllResponse, err error) {
	roleList := []types.Role{}
	param := &systemservice.Empty{}
	roles, err := l.svcCtx.SystemRpcClient.RoleAll(l.ctx, param)
	if err != nil {
		s, _ := status.FromError(err)
		if s.Message() == sql.ErrNoRows.Error() {
			return &types.RoleAllResponse{
				HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: "OK"},
				Total:              0,
				List:               roleList,
			}, nil
		}
		return nil, errorx.NewByCode(err, errorx.GRPC_ERROR).WithMeta("*SystemRpcClient.AllRoleList", err.Error(), param)
	}
	var (
		state = map[bool]string{
			true:  "deleted",
			false: "resume",
		}
	)
	for _, prole := range roles.Roles {
		role := types.Role{
			ID:         prole.ID,
			Role:       prole.Role,
			Name:       prole.Name,
			CreateBy:   prole.CreateBy,
			CreateTime: prole.CreateTime,
			UpdateBy:   prole.UpdateBy,
			UpdateTime: prole.UpdateTime,
			DeleteBy:   prole.DeleteBy,
			DeleteTime: prole.DeleteTime,
			State:      state[prole.DeleteTime != 0],
		}
		roleList = append(roleList, role)
	}

	return &types.RoleAllResponse{
		HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: "OK"},
		Total:              len(roleList),
		List:               roleList,
	}, nil
}
