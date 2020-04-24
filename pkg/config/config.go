package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Config interface {
	Load(p []string)
	InitConfig(c interface{}) (interface{}, error)
}

type CustomCfg interface {
	UpdateFun()
}

type BaseConfigOpt func(c *BaseConfig)


type BaseConfig struct {
	v          *viper.Viper
	isMonitor  bool
	customConf CustomCfg
}

func NewBaseConfig(opt ...BaseConfigOpt) *BaseConfig {
	b := &BaseConfig{
		v: viper.New(),
	}

	for _, c := range opt {
		c(b)
	}
	return b
}

// confPath map struct ----> key:filename value:file path
func (b *BaseConfig) Load(confPath map[string]string) {
	for fileName, filePath := range confPath {
		b.v.SetConfigName(fileName)
		b.v.AddConfigPath(filePath)
	}
	if err := b.v.ReadInConfig(); err != nil {
		log.Fatalf("Load ReadInConfig faile (%+v)", err)
		return
	}
	return
}

func (b *BaseConfig) StartMonitor() {
	if !b.isMonitor {
		log.Fatal("you must set monitor flag")
		return
	}
	const (
		debounceTimeout = time.Second * 5
	)

	// loop the config change
	for {
		b.v.WatchConfig()

		updateEven := make(chan *fsnotify.Event)

		go func() {
			var timer <-chan time.Time
			for {
				select {
				case <-timer:
					timer = nil
					if err := b.v.Unmarshal(b.customConf); err == nil {
						log.Println("begin update the config")
						b.customConf.UpdateFun()
					}

				case <-updateEven:
					if timer == nil {
						timer = time.After(debounceTimeout)
					}
				}
			}
		}()


		b.v.OnConfigChange(func(ev fsnotify.Event) {
			updateEven <- &ev
		})

		// 增加系统信号量，终止监听主goroutine
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		<-sigs
		return

	}
}

// s 是指指定的不同配置文件实例结构体
func (b *BaseConfig) InitConfig(s interface{}) (interface{}, error) {
	if err := b.v.Unmarshal(s); err != nil {
		log.Fatalf("config unmarshal faile (%+v)", err)
		return nil, err
	}
	return s, nil
}

// 获取指定的key值

func (b *BaseConfig) GetString(k string) string {
	return b.v.GetString(k)
}

func WithBaseConfigIsMonitor(m bool) BaseConfigOpt {
	return func(c *BaseConfig) {
		c.isMonitor = m
	}

}

func WithBaseConfigCustomConf(cfg CustomCfg) BaseConfigOpt {
	return func(c *BaseConfig) {
		c.customConf = cfg
	}
}

