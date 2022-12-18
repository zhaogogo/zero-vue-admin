package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/casbinx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/middleSvcCtx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/config"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/middleware"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"
)

type ServiceContext struct {
	Config          config.Config
	Casbin          rest.Middleware
	CheckUserExists rest.Middleware
	ParseJWTToken   rest.Middleware
	SystemRpcClient systemservice.SystemService
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlxConn := sqlx.NewMysql(c.Mysql.DataSource)
	sql, err := sqlxConn.RawDB()
	if err != nil {
		panic(err)
	}
	casbinx.Casbin.SetUp2(sql)
	//err := casbinx.Casbin.SetUp(c.Mysql.DataSource)

	svc := &ServiceContext{
		Config:          c,
		Casbin:          middleware.NewCasbinMiddleware().Handle,
		CheckUserExists: middleware.NewCheckUserExistsMiddleware().Handle,
		ParseJWTToken:   middleware.NewParseJWTTokenMiddleware().Handle,
		SystemRpcClient: systemservice.NewSystemService(zrpc.MustNewClient(c.SystemAdminRpcConf)),
	}

	middleSvcCtx.SetUp(svc.SystemRpcClient)

	return svc
}
