package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/model"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserPagingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserPagingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserPagingLogic {
	return &UserPagingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserPagingLogic) UserPaging(in *pb.UserPagingRequest) (*pb.UserPagingResponse, error) {
	var (
		err      error
		userList []model.User
	)
	if in.Page == 0 && in.PageSize == 0 {
		userList, err = l.svcCtx.UserModel.FindAll_NC(l.ctx)
		if err != nil {
			if err == sqlc.ErrNotFound {
				return nil, err
			}
			return nil, errors.Wrap(err, "数据库查询失败")
		}
	} else {
		userList, err = l.svcCtx.UserModel.FindPagingList_NC(l.ctx, &model.PagingUserList{Page: in.Page, PageSize: in.PageSize, NameX: in.NameX})
		if err != nil {
			if err == sqlc.ErrNotFound {
				return nil, err
			}
			return nil, errors.Wrap(err, "数据库查询失败")
		}
	}

	pUserList := []*pb.User{}
	for _, user := range userList {
		pUser := &pb.User{
			ID:         user.Id,
			Name:       user.Name,
			NickName:   user.NickName,
			PassWord:   user.Password,
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
		}

		if user.DeleteTime.Valid {
			pUser.DeleteTime = user.DeleteTime.Time.Unix()
		} else {
			pUser.DeleteTime = 0
		}

		pUserList = append(pUserList, pUser)
	}

	return &pb.UserPagingResponse{
		Users: pUserList,
	}, nil
}
