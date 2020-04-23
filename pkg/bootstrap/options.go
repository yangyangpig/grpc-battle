package bootstrap

import "google.golang.org/grpc"

type serverOptions struct {
	logLevel string

	grpcAddr string

	httpAddr string

	grpcServer     func(s *grpc.Server)
	grpcServerOpts []grpc.ServerOption

	ownerConfig struct {
		config   interface{}
		onUpdate func()
	}
}
