package bootstrap

import (
	"fmt"
	"github.com/spf13/cobra"
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
	serverOptions serverOptions
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
			//b.conf.Load(map[string]string {
			//	b.name: ".",
			//	b.name: "$HOME/." + b.name,
			//	b.name: "/etc",
			//	b.name: "$HOME/etc",
			//})

			// 服务的启动
			if err := b.mixtureServer.InitServer(); err != nil {
				log.Println(err)
				return err
			}

			b.mixtureServer.StartServer()


			gracefulClose := make(chan struct{})
			go func() {
				// waite the world
				b.mixtureServer.WaiteAllWorld()
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

	startCmd.PersistentFlags().StringVar(&b.serverOptions.grpcAddr, "listen", ":8080", "grpc address")
	startCmd.PersistentFlags().StringVar(&b.serverOptions.httpAddr, "admin", ":80", "admin address")
	startCmd.PersistentFlags().StringVar(&b.serverOptions.logLevel, "loglev", "info", "log level")
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
