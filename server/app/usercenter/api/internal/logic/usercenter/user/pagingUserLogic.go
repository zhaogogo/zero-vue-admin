package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	errorx2 "github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/usercenter/api/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/usercenter/api/internal/types"
)

type PagingUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPagingUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PagingUserLogic {
	return &PagingUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PagingUserLogic) PagingUser(req *types.PagingUserRequest) (resp *types.PagingUserResponse, err error) {
	users, err := l.svcCtx.UserModel.FindPaging(l.ctx, req.Page, req.PageSize)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return &types.PagingUserResponse{
				HttpCommonResponse:  types.HttpCommonResponse{Code: 200, Msg: err.Error()},
				TableCommonResponse: types.TableCommonResponse{PageSize: req.PageSize, Page: req.Page, Total: len(users)},
				List:                nil,
			}, nil
		}
		return nil, errorx2.New(errorx2.SERVER_COMMON_ERROR, "获取分页用户失败")
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
	return &types.PagingUserResponse{
		HttpCommonResponse:  types.HttpCommonResponse{Code: 200, Msg: "OK"},
		TableCommonResponse: types.TableCommonResponse{PageSize: req.PageSize, Page: req.Page},
		List:                list,
	}, nil
}
