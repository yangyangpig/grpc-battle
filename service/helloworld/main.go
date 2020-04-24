package main

import (
	"google.golang.org/grpc"
	"grpc-battle/pkg/bootstrap"
	"grpc-battle/pkg/config"
	pb "grpc-battle/protocol/helloWorld"
	"grpc-battle/service/helloworld/server"
)
var serverName = "hello-world"
func main()  {
	// 加载配置
	svrCfg := config.NewBaseConfig(config.WithBaseConfigIsMonitor(false))
	svrCfg.Load(map[string]string {
		serverName: ".",
		serverName: "$HOME/." + serverName,
		serverName: "/etc",
		serverName: "$HOME/etc",
		})

	// grpc服务注册
	helloWorld := bootstrap.NewGrpcServer (
		bootstrap.WithGrpcServerRegister(func(s *grpc.Server) {
			h := server.NewHelloWorld()
			pb.RegisterHelloWorldServer(s, h)
		}),
		bootstrap.WithGrpcServerBaseConf(svrCfg),
		)

	bootstrap.New(
		bootstrap.WithBootstrapName(serverName),
		bootstrap.WithBootstrapMixtureServer(helloWorld),
		).Start()
}
