package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/esManager/internal/config"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/esManager/internal/server"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/esManager/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/rpc/esManager/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/esmanager.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterEsManagerServiceServer(grpcServer, server.NewEsManagerServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	logx.DisableStat()
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
