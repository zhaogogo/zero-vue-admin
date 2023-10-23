package alarm

import (
	"context"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WebhookLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWebhookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WebhookLogic {
	return &WebhookLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WebhookLogic) Webhook(req *types.AlarmRequest) error {
	l.svcCtx.SlienceList.Mu.RLock()
	matchDefault := l.svcCtx.SlienceList.Sliences
	defer l.svcCtx.SlienceList.Mu.RUnlock()

	for _, alert := range req.Alerts {
		AlarmIsMatchDefault(alert, l.svcCtx)
	}
	return nil
}

func AlarmIsMatchDefault(alert types.Alerts, svcCtx *svc.ServiceContext) {
	isMatch := true
	for instance, m := range svcCtx.SlienceList.Sliences {

	}
}
