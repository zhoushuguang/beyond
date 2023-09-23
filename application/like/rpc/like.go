package main

import (
	"flag"
	"fmt"

	"beyond/application/like/rpc/internal/config"
	"beyond/application/like/rpc/internal/server"
	"beyond/application/like/rpc/internal/svc"
	"beyond/application/like/rpc/service"

	"github.com/zeromicro/go-zero/core/conf"
	zs "github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/like.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		service.RegisterLikeServer(grpcServer, server.NewLikeServer(ctx))

		if c.Mode == zs.DevMode || c.Mode == zs.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
