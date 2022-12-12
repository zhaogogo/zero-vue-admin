package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/usercenter/api/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/usercenter/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AllUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAllUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AllUserLogic {
	return &AllUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AllUserLogic) AllUser() (resp *types.AllUserResponse, err error) {
	users, err := l.svcCtx.UserModel.FindAll(l.ctx)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return &types.AllUserResponse{
				HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: err.Error()},
				Total:              0,
				List:               nil,
			}, nil
		}
	}
	state := map[bool]string{
		true:  "deleted",
		false: "resume",
	}
	list := []types.User{}
	for _, user := range users {
		ut := types.User{
			ID:         user.Id,
			Name:       user.Name,
			NickName:   user.NickName,
			PassWord:   "****",
			UserType:   user.Type,
			Email:      user.Email,
			Phone:      user.Phone,
			Department: user.Department,
			Position:   user.Position,
			CreateBy:   user.CreateBy,
			CreateTime: user.CreateTime.Unix(),
			UpdateBy:   user.UpdateBy,
			UpdateTime: user.UpdateTime.Unix(),
			DeleteBy:   user.DeleteBy,
			State:      state[user.DeleteTime.Valid],
		}
		if user.DeleteTime.Valid {
			ut.DeleteTime = user.DeleteTime.Time.Unix()
		} else {
			ut.DeleteTime = 0
		}
		list = append(list, ut)
	}

	return &types.AllUserResponse{
		HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: "OK"},
		Total:              len(list),
		List:               list,
	}, nil
}
