package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PagingUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPagingUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PagingUserListLogic {
	return &PagingUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PagingUserListLogic) PagingUserList(in *pb.PagingRequest) (*pb.PagingUserListResponse, error) {
	userList, err := l.svcCtx.UserModel.FindListPaging(l.ctx, in.Page, in.PageSize)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return &pb.PagingUserListResponse{List: []*pb.User{}}, nil
		}
		return nil, errors.Wrapf(err, "数据库分页查询失败, 表: user, page=%d, pageSize=%d", in.Page, in.PageSize)
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

	return &pb.PagingUserListResponse{
		List: pUserList,
	}, nil
}
