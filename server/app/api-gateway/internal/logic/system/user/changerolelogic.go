package user

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeRoleLogic {
	return &ChangeRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeRoleLogic) ChangeRole(req *types.UserChangeRoleRequest) (resp *types.HttpCommonResponse, err error) {
	hasRolein := false
	userid := l.ctx.Value("user_id").(uint64)
	userroleinfo := l.ctx.Value("userroleinfo").([]*systemservice.UserRole)
	for _, userrole := range userroleinfo {
		if userrole.RoleID == req.RoleID {
			hasRolein = true
			break
		}
	}
	if !hasRolein {
		return nil, errorx.New(errors.New("未分配的角色,刷新页面重试"), "未分配的角色,刷新页面重试")
	}
	updateUserCurrentRoleParam := &systemservice.UpdateUserCurrentRoleRequest{UserID: userid, RoleID: req.RoleID}
	_, err = l.svcCtx.SystemRpcClient.UpdateUserCurrentRole(l.ctx, updateUserCurrentRoleParam)
	if err != nil {
		return nil, errorx.New(err, "切换用户角色失败").WithMeta("SystemRpcClient.UpdateUserCurrentRole", err.Error(), updateUserCurrentRoleParam)
	}
	return &types.HttpCommonResponse{Code: 200, Msg: "OK"}, nil
}
