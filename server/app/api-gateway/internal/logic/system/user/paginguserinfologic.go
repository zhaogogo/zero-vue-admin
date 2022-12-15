package user

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"
	"google.golang.org/grpc/status"
	"strings"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PagingUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPagingUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PagingUserInfoLogic {
	return &PagingUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PagingUserInfoLogic) PagingUserInfo(req *types.PagingUserRequest) (resp *types.PagingUserResponse, err error) {
	userlist, err := l.svcCtx.SystemRpcClient.PagingUserList(l.ctx, &systemservice.PagingRequest{Page: req.Page, PageSize: req.PageSize})
	if err != nil {
		e, _ := status.FromError(err)
		if e.Message() == sql.ErrNoRows.Error() {
			return &types.PagingUserResponse{
				HttpCommonResponse:   types.HttpCommonResponse{Code: 200, Msg: "OK"},
				PagingCommonResponse: types.PagingCommonResponse{Page: req.Page, PageSize: req.PageSize, Total: userlist.Total},
				List:                 []types.User{},
			}, nil
		}
		return nil, err
	}
	var (
		state = map[bool]string{
			true:  "deleted",
			false: "resume",
		}
		msgErrList = errorx.MsgErrList{}
		msg        = "OK"
	)
	list := []types.User{}
	for _, user := range userlist.List {
		ut := types.User{
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
			UpdateBy:   user.UpdateBy,
			UpdateTime: user.UpdateTime,
			DeleteBy:   user.DeleteBy,
			DeleteTime: user.DeleteTime,
			State:      state[user.DeleteTime != 0],
		}
		list = append(list, ut)
	}

	//获取用户角色
	mr.MapReduce(
		func(source chan<- interface{}) {
			for _, user := range list {
				source <- user.ID
			}
		},
		func(item interface{}, writer mr.Writer, cancel func(error)) {
			userid := item.(uint64)
			userroles, err := l.svcCtx.SystemRpcClient.GetUserRoleByUserID(l.ctx, &systemservice.UserID{ID: userid})
			if err != nil {
				l.Error(err)
				msgErrList.Append(fmt.Sprintf("获取用户角色错误 [user_id: %d, error: %v]", userid, err.Error()))
				return
			}

			writer.Write(userroles)
		},
		func(pipe <-chan interface{}, writer mr.Writer, cancel func(error)) {
			for v := range pipe {
				userroles, ok := v.(*systemservice.UserRoleList)
				if ok {
					for _, userrole := range userroles.UserRole {
						for index, user := range list {
							if user.ID == userrole.UserID {
								list[index].RoleList = append(list[index].RoleList, userrole.RoleID)
							}
						}
					}
				} else {
					logx.Errorf("mr reducer断言失败, 实际类型: (%T), 断言类型: (*systemservice.UserRoleList)", v, v)
				}

			}
		},
	)
	if len(msgErrList.List) != 0 {
		msg = strings.Join(msgErrList.List, " | ")
	}
	return &types.PagingUserResponse{
		HttpCommonResponse:   types.HttpCommonResponse{Code: 200, Msg: msg},
		PagingCommonResponse: types.PagingCommonResponse{Page: req.Page, PageSize: req.PageSize, Total: int64(len(list))},
		List:                 list,
	}, nil
}
