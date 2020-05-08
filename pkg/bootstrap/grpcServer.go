package bootstrap

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"grpc-battle/pkg/config"
	"log"
	"net"
	"sync"
)

type GrpcServerOpt func(g *GrpcServer)

type GrpcServer struct {
	grpcAddr           string
	logLevel           string
	grpcServer         *grpc.Server
	grpcServerRegister func(s *grpc.Server)
	grpcServerOpt      []grpc.ServerOption
	GRPCListener       net.Listener
	wg                 sync.WaitGroup
	conf               *config.BaseConfig
}

func WithGrpcServerRegister(gsb func(s *grpc.Server)) GrpcServerOpt {
	return func(g *GrpcServer) {
		g.grpcServerRegister = gsb
	}
}

func WithGrpcServerOpt(gop ...grpc.ServerOption) GrpcServerOpt {
	return func(g *GrpcServer) {
		if g.grpcServerOpt == nil {
			g.grpcServerOpt = gop
		} else {
			g.grpcServerOpt = append(g.grpcServerOpt, gop...)
		}
	}
}

func WithGrpcServerBaseConf(f *config.BaseConfig) GrpcServerOpt {
	return func(g *GrpcServer) {
		g.conf = f
		g.logLevel = g.conf.GetString("logLevel")
		g.grpcAddr = g.conf.GetString("listen")
	}
}

func NewGrpcServer(opts ...GrpcServerOpt) *GrpcServer {
	if len(opts) == 0 {
		// set default opt
		return &GrpcServer{}
	}

	g := &GrpcServer{}

	for _, opt := range opts {
		opt(g)
	}
	log.Printf("grpc server init success (%+v)", g)

	return g

}

func (g *GrpcServer) InitServer() error {
	// new grpc server
	g.grpcServer = grpc.NewServer(g.grpcServerOpt...)
	// register customer server
	g.grpcServerRegister(g.grpcServer)

	// reflection
	reflection.Register(g.grpcServer)

	// server health register
	hsvr := health.NewServer()
	hsvr.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(g.grpcServer, hsvr)

	log.Printf("grpc address (%s)", g.grpcAddr)
	log.Printf("grpc loglevel (%s)", g.logLevel)
	var err error
	g.GRPCListener, err = net.Listen("tcp", g.grpcAddr)
	if err != nil {
		return err
	}

	return err
}

func (g *GrpcServer) StartServer() {
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		err := g.grpcServer.Serve(g.GRPCListener)
		log.Printf("start server done: %v", err)
	}()
}

func (g *GrpcServer) CloseServer() {
	g.grpcServer.GracefulStop()
}
// 这个wait是阻塞等待当前的协程
func (g *GrpcServer) WaitAllWorld() {
	g.wg.Wait()
}

func (g *GrpcServer) CleanAllServer(s chan struct{}) {
	go func() {
		<-s
		g.grpcServer.GracefulStop()

	}()

}
