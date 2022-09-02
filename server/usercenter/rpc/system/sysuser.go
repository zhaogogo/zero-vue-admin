package main

import (
	"flag"
	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/rpc/system/pkg/interceptors"

	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/rpc/system/internal/config"
	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/rpc/system/internal/server"
	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/rpc/system/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/sysuser.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterSystemUserServer(grpcServer, server.NewSystemUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	s.AddUnaryInterceptors(interceptors.GrpcLogAndValidate)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
