package svc

import (
	"bytes"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/config"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/middleware"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/casbinx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/gormc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/middleSvcCtx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/esManager/esmanagerservice"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/monitoring/monitoringmanager"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"html/template"
	"sync"
)

type Use func() *gorm.DB

type ServiceContext struct {
	Config      config.Config
	SlienceList SafeSliences

	Casbin          rest.Middleware
	CheckUserExists rest.Middleware
	ParseJWTToken   rest.Middleware

	DB           *gorm.DB
	MonitoringDB Use

	SystemRpcClient    systemservice.SystemService
	ESManagerRpcClient esmanagerservice.EsManagerService
	MonitoringRpcConf  monitoringmanager.MonitoringManager
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlxConn := sqlx.NewMysql(c.Mysql.System.DataSource)
	sql, err := sqlxConn.RawDB()
	if err != nil {
		panic(err)
	}
	casbinx.Casbin.SetUp2(sql)
	//err := casbinx.Casbin.SetUp(c.Mysql.DataSource)
	db := gormc.MustNewGrom(c.Mysql)
	db = db.Debug()
	svc := &ServiceContext{
		Config:          c,
		Casbin:          middleware.NewCasbinMiddleware().Handle,
		CheckUserExists: middleware.NewCheckUserExistsMiddleware().Handle,
		ParseJWTToken:   middleware.NewParseJWTTokenMiddleware().Handle,
		DB:              db,
		MonitoringDB:    func() *gorm.DB { return db.Clauses(dbresolver.Use("monitoring")) },
		//MonitoringDB:       db.Clauses(dbresolver.Use("monitoring")),

		SystemRpcClient:    systemservice.NewSystemService(zrpc.MustNewClient(c.SystemAdminRpcConf)),
		ESManagerRpcClient: esmanagerservice.NewEsManagerService(zrpc.MustNewClient(c.ESManagerRpcConf)),
		MonitoringRpcConf:  monitoringmanager.NewMonitoringManager(zrpc.MustNewClient(c.MonitoringRpcConf)),
	}
	middleSvcCtx.SetUp(svc.SystemRpcClient)
	svc.SlienceList = GetSliences(svc)
	return svc
}

func GetSliences(svc *ServiceContext) SafeSliences {
	slienceResults := []types.SlienceJoinRest{}
	svc.MonitoringDB().Model(&types.Host{}).
		Select("`hosts`.`id`,`hosts`.`host`, `slience_names`.`default`,`slience_names`.`slience_name`, `slience_matchers`.`name`, `slience_matchers`.`value`, `slience_matchers`.`is_regex`, `slience_matchers`.`is_equal`").
		Joins("JOIN slience_names ON hosts.id = slience_names.host_id").
		Joins("JOIN slience_matchers ON slience_names.id = slience_matchers.slience_name_id and and slience_matchers.host_id = hosts.id").
		Scan(&slienceResults)

	templateValue := bytes.NewBuffer(nil)
	for i, slienceResult := range slienceResults {
		parse, err := template.New("t1").Parse(slienceResult.Value)
		if err != nil {
			logx.Error(err)
			continue
		}
		err = parse.ExecuteTemplate(templateValue, "t1", parse)
		if err != nil {
			logx.Error(err)
			continue
		}
		slienceResults[i].Value = templateValue.String()
		templateValue.Reset()
	}
	sliences := SafeSliences{Sliences: make(map[string]map[string][]types.Matchers)}
	sliences.Mu.Lock()
	defer sliences.Mu.Unlock()
	for _, res := range slienceResults {
		if sliences.Sliences[res.Host] == nil {
			sliences.Sliences[res.Host] = make(map[string][]types.Matchers)
		}
		if res.Default {
			sliences.Sliences[res.Host]["default"] = append(sliences.Sliences[res.Host]["default"], types.Matchers{
				Name:    res.Name,
				Value:   res.Value,
				IsRegex: false,
				IsEqual: true,
			})
		}
		sliences.Sliences[res.Host][res.SlienceName] = append(sliences.Sliences[res.Host][res.SlienceName], types.Matchers{
			Name:    res.Name,
			Value:   res.Value,
			IsRegex: false,
			IsEqual: true,
		})
	}
	return sliences
}

type SafeSliences struct {
	Mu sync.RWMutex
	//           instance   alertName
	Sliences map[string]map[string][]types.Matchers
}
