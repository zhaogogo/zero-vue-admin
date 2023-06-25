package role

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"
	"google.golang.org/grpc/status"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.RoleDetailRequest) (resp *types.RoleDetailResponse, err error) {
	param := &systemservice.RoleID{ID: req.ID}
	role, err := l.svcCtx.SystemRpcClient.RoleDetail(l.ctx, param)
	if err != nil {
		s, _ := status.FromError(err)
		if s.Message() == sql.ErrNoRows.Error() {
			return nil, errorx.New(err, fmt.Sprintf("获取角色详情失败,id=%v无数据", req.ID)).WithMeta("SystemRpcClient.RoleDetail", err.Error(), param)
		}
		return nil, errorx.New(err, "获取角色详情失败").WithMeta("SystemRpcClient.RoleDetail", err.Error(), param)
	}

	return &types.RoleDetailResponse{
		HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: "OK"},
		Detail: types.Role{
			ID:         role.ID,
			Role:       role.Role,
			Name:       role.Name,
			CreateBy:   role.CreateBy,
			CreateTime: role.CreateTime,
			UpdateBy:   role.UpdateBy,
			UpdateTime: role.UpdateTime,
			DeleteBy:   role.DeleteBy,
			DeleteTime: role.DeleteTime,
		},
	}, nil
}
