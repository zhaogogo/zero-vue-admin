package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/config"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/handler"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/responseerror"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/apigateway-api.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/",
		Handler: http.FileServer(http.Dir("assets")).ServeHTTP,
	})

	httpx.SetErrorHandler(responseerror.ErrorHandle)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
