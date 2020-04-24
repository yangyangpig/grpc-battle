package bootstrap

import (
	"fmt"
	"github.com/spf13/cobra"
	"grpc-battle/pkg/config"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type BootstrapOpt func(b *Bootstrap)

type Bootstrap struct {
	name          string
	rootCmd       *cobra.Command
	mixtureServer MixtureServer
}

func WithBootstrapName(name string) BootstrapOpt {
	return func(b *Bootstrap) {
		b.name = name
	}
}


func WithBootstrapMixtureServer(s MixtureServer) BootstrapOpt {
	return func(b *Bootstrap) {
		b.mixtureServer = s
	}
}


func New(opts ...BootstrapOpt) *Bootstrap {
	b := &Bootstrap{
		rootCmd: &cobra.Command{},
	}
	// init the bootstrap opt
	for _, opt := range opts {
		opt(b)
	}

	b.init()
	return b
}


func (b *Bootstrap) init() {
	startCmd := &cobra.Command {
		Use:   "start",
		Short: fmt.Sprintf("start the server name %s", b.name),
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := b.mixtureServer.InitServer(); err != nil {
				log.Println(err)
				return err
			}

			b.mixtureServer.StartServer()


			gracefulClose := make(chan struct{})
			go func() {
				// wait the world
				b.mixtureServer.WaitAllWorld()
				close(gracefulClose)
			}()

			stop := make(chan struct{})
			// 系统信号关闭
			waitSignal(stop)
			// 服务优雅关闭
			b.mixtureServer.CleanAllServer(stop)

			select {
			case <-gracefulClose:
				log.Println("graceful stopped server")
			case <-time.After(time.Second * 3):
				log.Println("graceful stopped server timeout")

			}
			return nil
		},
	}

	// require the server config
	var baseConfig *config.BaseConfig
	if grpc, ok := b.mixtureServer.(*GrpcServer); ok {
		baseConfig = grpc.conf
	}



	listen := baseConfig.GetString("listen")
	httpAddr := baseConfig.GetString("admin")
	logLevel := baseConfig.GetString("loglev")

	startCmd.PersistentFlags().StringVar(&listen, "listen", ":8080", "grpc address")
	startCmd.PersistentFlags().StringVar(&httpAddr, "admin", ":80", "admin address")
	startCmd.PersistentFlags().StringVar(&logLevel, "loglev", "info", "log level")
	b.rootCmd.AddCommand(startCmd)
}

func (b *Bootstrap) Start() error {
	return b.rootCmd.Execute()
}

func waitSignal(s chan struct{})  {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	close(s)

}
