package middleSvcCtx

import (
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"
)

var Ctx *middleSvcCtx = new(middleSvcCtx)

type middleSvcCtx struct {
	SystemRpcClient systemservice.SystemService
}

func (m *middleSvcCtx) SetUp(systemRpcClient systemservice.SystemService) {
	m.SystemRpcClient = systemRpcClient
}

func SetUp(systemRpcClient systemservice.SystemService) {
	Ctx.SetUp(systemRpcClient)
}
