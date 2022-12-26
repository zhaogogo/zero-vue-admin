package user

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

func (l *DetailLogic) Detail(req *types.UserDetailRequest) (resp *types.UserDetailResponse, err error) {
	username := l.ctx.Value("userName").(string)
	userDetailParam := &systemservice.UserID{ID: req.ID}
	user, err := l.svcCtx.SystemRpcClient.UserDetail(l.ctx, userDetailParam)
	if err != nil {
		s, _ := status.FromError(err)
		if s.Message() == sql.ErrNoRows.Error() {
			return nil, errorx.NewByCode(err, errorx.DB_NOTFOUND).WithMeta("SystemRpcClient.UserDetail", err.Error(), userDetailParam)
		}
		return nil, errorx.NewByCode(err, errorx.GRPC_ERROR).WithMeta("SystemRpcClient.UserDetail", err.Error(), userDetailParam)
	}
	var state = map[bool]string{
		true:  "deleted",
		false: "resume",
	}
	u := types.User{
		ID:         user.ID,
		Name:       user.Name,
		NickName:   user.NickName,
		PassWord:   user.PassWord,
		UserType:   user.UserType,
		Email:      user.Email,
		Phone:      user.Phone,
		Department: user.Department,
		Position:   user.Position,
		CreateBy:   user.CreateBy,
		CreateTime: user.CreateTime,
		UpdateBy:   username,
		UpdateTime: user.UpdateTime,
		DeleteBy:   user.DeleteBy,
		DeleteTime: user.DeleteTime,
		State:      state[user.DeleteTime != 0],
	}
	return &types.UserDetailResponse{
		HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: "OK"},
		Detail:             u,
	}, err
}
