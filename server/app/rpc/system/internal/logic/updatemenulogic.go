package logic

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/model/system"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuLogic {
	return &UpdateMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMenuLogic) UpdateMenu(in *pb.UpdateMenuRequest) (*pb.Empty, error) {
	p := &system.Menu{
		Id:        in.Id,
		ParentId:  in.ParentID,
		Name:      in.Name,
		Path:      in.Path,
		Component: in.Component,
		Title:     in.Title,
		Icon:      in.Icon,
		Sort:      in.Sort,
		Hidden:    in.Hiddent,
	}
	err := l.svcCtx.MenuModel.Update(l.ctx, p)
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
