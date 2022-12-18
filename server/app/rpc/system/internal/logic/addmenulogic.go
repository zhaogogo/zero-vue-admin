package logic

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/model/system"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMenuLogic {
	return &AddMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddMenuLogic) AddMenu(in *pb.AddMenuRequest) (*pb.Empty, error) {
	_, err := l.svcCtx.MenuModel.Insert(l.ctx, &system.Menu{
		Id:        0,
		ParentId:  in.ParentID,
		Name:      in.Name,
		Path:      in.Path,
		Component: in.Component,
		Title:     in.Title,
		Icon:      in.Icon,
		Sort:      in.Sort,
		Hidden:    in.Hiddent,
	})
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
