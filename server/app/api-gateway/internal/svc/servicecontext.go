package svc

import (
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/config"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/middleware"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/casbinx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/gormc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/middleSvcCtx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/slience"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/esManager/esmanagerservice"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/monitoring/monitoringmanager"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/system/systemservice"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"html/template"
	"path/filepath"
	"strconv"
	"strings"
)

type Use func() *gorm.DB

type ServiceContext struct {
	Config         config.Config
	SlienceList    *slience.SafeSliences
	NotifyTemplate *template.Template

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
	setUpZerolog(c.Log)
	zlog.Info().Msg("start")
	sqlxConn := sqlx.NewMysql(c.Mysql.System.DataSource)
	sql, err := sqlxConn.RawDB()
	if err != nil {
		panic(err)
	}
	casbinx.Casbin.SetUp2(sql)
	//err := casbinx.Casbin.SetUp(c.Mysql.DataSource)
	db := gormc.MustNewGrom(c.Mysql)
	db = db.Debug()
	notifyTemplate, err := template.New("notify").ParseFiles(c.MonitoringConfig.NotifyTemplatePath)
	if err != nil {
		panic(err)
	}

	svc := &ServiceContext{
		Config:          c,
		SlienceList:     &slience.SafeSliences{},
		NotifyTemplate:  notifyTemplate,
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
	err = slience.GetConsumerSliences(svc.MonitoringDB(), svc.SlienceList)
	return svc
}

func setUpZerolog(c logx.LogConf) {
	f := &lumberjack.Logger{
		Filename:   filepath.Join(c.Path, "service.log"),
		MaxSize:    c.MaxSize,
		MaxAge:     0,
		MaxBackups: c.MaxBackups,
		LocalTime:  false,
		Compress:   false,
	}
	log := zerolog.New(f).With().Timestamp().CallerWithSkipFrameCount(2).Logger()
	zlog.Logger = log
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return strings.Replace(file, "C:/Users/qiang.zhao/Desktop/owner/zero-vue-admin/server/app/", "", 1) + ":" + strconv.Itoa(line)
	}
	switch c.Level {
	case "info":
		zlog.Logger.Level(zerolog.DebugLevel)
	case "error":
		zlog.Logger.Level(zerolog.ErrorLevel)
	case "severe":
		zlog.Logger.Level(zerolog.DebugLevel)
	}
}
