package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/model/system"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserMenuParamsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserMenuParamsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserMenuParamsLogic {
	return &UpdateUserMenuParamsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserMenuParamsLogic) UpdateUserMenuParams(in *pb.UpdateUserMenuParamsRequest) (*pb.Empty, error) {
	insertData := []*system.UserMenuParams{}
	updateData := []*system.UserMenuParams{}
	for _, usermenuparam := range in.UserMenuParams {
		if usermenuparam.ID == 0 {
			insertData = append(insertData, &system.UserMenuParams{
				Id:     usermenuparam.ID,
				UserId: usermenuparam.UserID,
				MenuId: in.MenuId,
				Type:   usermenuparam.Type,
				Key:    usermenuparam.Key,
				Value:  usermenuparam.Value,
			})
		} else {
			updateData = append(updateData, &system.UserMenuParams{
				Id:     usermenuparam.ID,
				UserId: usermenuparam.UserID,
				MenuId: in.MenuId,
				Type:   usermenuparam.Type,
				Key:    usermenuparam.Key,
				Value:  usermenuparam.Value,
			})
		}
	}
	if len(insertData) == 0 && len(updateData) == 0 {
		err := l.svcCtx.UserMenuParamsModel.DeleteByMenuID(l.ctx, in.MenuId)
		if err != nil {
			return nil, err
		}
		return &pb.Empty{}, nil
	}

	if len(insertData) != 0 && len(updateData) == 0 {
		_, err := l.svcCtx.UserMenuParamsModel.InsertMultiple(l.ctx, insertData)
		if err != nil {
			return nil, err
		}
		return &pb.Empty{}, nil
	}

	if len(insertData) == 0 && len(updateData) != 0 {
		err := l.svcCtx.UserMenuParamsModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
			if err := l.svcCtx.UserMenuParamsModel.TransDeleteNotINANDMenuID(ctx, session, updateData, in.MenuId); err != nil {
				return err
			}
			for _, d := range updateData {
				err := l.svcCtx.UserMenuParamsModel.TransUpdate(ctx, session, &system.UserMenuParams{
					Id:     d.Id,
					UserId: d.UserId,
					MenuId: d.MenuId,
					Type:   d.Type,
					Key:    d.Key,
					Value:  d.Value,
				})
				if err != nil {
					return err
				}
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
		return &pb.Empty{}, nil
	}
	if len(insertData) != 0 && len(updateData) != 0 {
		err := l.svcCtx.UserMenuParamsModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
			if err := l.svcCtx.UserMenuParamsModel.TransDeleteNotINANDMenuID(ctx, session, updateData, updateData[0].MenuId); err != nil {
				return err
			}
			for _, d := range updateData {
				err := l.svcCtx.UserMenuParamsModel.TransUpdate(ctx, session, d)
				if err != nil {
					return err
				}
			}

			res, err := l.svcCtx.UserMenuParamsModel.TransInsertMultiple(ctx, session, insertData)
			if err != nil {
				return err
			}
			a, err := res.RowsAffected()
			if err != nil {
				return err
			}
			if a != int64(len(insertData)) {
				return errors.New(fmt.Sprintf("数据库插入行数(%v)与提交参数(%v)不匹配", a, len(insertData)))
			}

			return nil
		})
		if err != nil {
			return nil, err
		}
		return &pb.Empty{}, nil
	}

	return nil, errors.New(fmt.Sprintf("insert data count (%v), update data count(%v)", len(insertData), len(updateData)))
}
