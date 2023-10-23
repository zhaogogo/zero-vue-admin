package alertrule

import (
	"context"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/model/monitoring"

	"github.com/zeromicro/go-zero/core/logx"
)

type PagingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPagingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PagingLogic {
	return &PagingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PagingLogic) Paging(req *types.AlertRulePagingRequest) (resp *types.AlertRuleResponse, err error) {
	alertRulesData := []monitoring.AlertRules{}
	var alertRulesCount int64
	db := l.svcCtx.MonitoringDB()
	if err = db.Model(alertRulesData).Count(&alertRulesCount).Error; err != nil {
		return nil, errorx.New(err, "AlertRules Count Failed")
	}
	err = db.Model(alertRulesData).Limit(int(req.PageSize)).Offset(int((req.Page - 1) * req.PageSize)).Preload("Labels").Preload("Tags").Preload("Querys").Find(&alertRulesData).Error
	if err != nil {
		return nil, errorx.New(err, "查找AlertRules失败")
	}
	data := []types.AlertRule{}
	for _, alertrule := range alertRulesData {
		var (
			deleteTime uint64
			labels     []types.AlertRuleLabel
			tags       []types.AlertRuleTag
			querys     []types.AlertRuleQuery
		)
		if alertrule.DeletedAt.Valid == true {
			deleteTime = 0
		} else {
			deleteTime = uint64(alertrule.DeletedAt.Time.Second())
		}
		for _, d := range alertrule.Labels {
			labels = append(labels, types.AlertRuleLabel{
				Id:          d.Id,
				AlertRuleId: d.AlertRuleId,
				Key:         d.Key,
				Value:       d.Value,
			})
		}
		for _, d := range alertrule.Tags {
			tags = append(tags, types.AlertRuleTag{
				Id:          d.Id,
				AlertRuleId: d.AlertRuleId,
				Key:         d.Key,
				Value:       d.Value,
			})
		}
		for _, d := range alertrule.Querys {
			querys = append(querys, types.AlertRuleQuery{
				Id:          d.Id,
				AlertRuleId: d.AlertRuleId,
				Query:       d.Query,
			})
		}

		data = append(data, types.AlertRule{
			Id:         uint64(alertrule.ID),
			Ttype:      alertrule.Ttype,
			Name:       alertrule.Name,
			Group:      alertrule.Group,
			To:         alertrule.To,
			Expr:       alertrule.Expr,
			For:        alertrule.For.String(),
			AlertText:  alertrule.AlertText,
			CreatedAt:  uint64(alertrule.CreatedAt.Second()),
			ModifiedAt: uint64(alertrule.ModifiedAt.Second()),
			DeletedAt:  deleteTime,
			Labels:     labels,
			Tags:       tags,
			Querys:     querys,
		})
	}
	return &types.AlertRuleResponse{
		HttpCommonResponse:   types.HttpCommonResponse{Code: 200, Msg: "OK", Meta: nil},
		PagingCommonResponse: types.PagingCommonResponse{Page: req.Page, PageSize: req.PageSize, Total: alertRulesCount},
		List:                 data,
	}, nil
}
