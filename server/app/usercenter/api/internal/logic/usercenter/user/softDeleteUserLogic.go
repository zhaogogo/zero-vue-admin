package user

import (
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	errorx2 "github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/usercenter/api/internal/pkg/logiccommon"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/usercenter/api/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/usercenter/api/internal/types"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type SoftDeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSoftDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SoftDeleteUserLogic {
	return &SoftDeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SoftDeleteUserLogic) SoftDeleteUser(req *types.SoftDeleteUserRequest) (resp *types.HttpCommonResponse, err error) {
	//获取删除用户的ID
	userid, err := logiccommon.GetUserIdWithJWT(l.ctx)
	if err != nil {
		return nil, errorx2.New(errorx2.SERVER_COMMON_ERROR, "获取用户ID失败")
	}
	user, err := l.svcCtx.UserModel.FindOneByIDWHEREDeleteTimeISNULL(l.ctx, userid)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return nil, errorx2.NewFromCode(errorx2.USER_NOT_FOUND, err)
		}
		return nil, errorx2.New(errorx2.SERVER_COMMON_ERROR, "获取用户失败")
	}
	switch req.State {
	case "deleted":
		if err := l.svcCtx.UserModel.UpdateDeleteColumn(l.ctx, req.ID, user.Name, sql.NullTime{Time: time.Now(), Valid: true}); err != nil {
			return nil, errorx2.NewFromCode(errorx2.DB_ERROR, err)
		}
	case "resume":
		if err := l.svcCtx.UserModel.UpdateDeleteColumn(l.ctx, req.ID, "", sql.NullTime{}); err != nil {
			return nil, errorx2.NewFromCode(errorx2.DB_ERROR, err)
		}
	}

	return &types.HttpCommonResponse{Code: 200, Msg: "OK"}, nil
}
