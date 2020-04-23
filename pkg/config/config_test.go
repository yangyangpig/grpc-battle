package config

import (
	"fmt"
	"grpc-battle/pkg/config/redis"
	"testing"
)

func TestNewBaseConfig(t *testing.T) {
	baseCfg := NewBaseConfig(WithBaseConfigIsMonitor(true))
	l := make(map[string]string)
	l["test"] = "/Users/liufuhong/github-program/grpc-battle/pkg/config"
	baseCfg.Load(l)
	rediscfg, err := baseCfg.InitConfig(&redis.RedisConfig{})
	if err != nil {
		t.Errorf("init config error (+%v)", err)
		return
	}
	// 使用时候，需要类型判断
	if f, ok := rediscfg.(*redis.RedisConfig); ok {
		fmt.Println(f.Print())
	}
	fmt.Println("-----------run end----------")
}

func TestBaseConfig_StartMonitor(t *testing.T) {
	baseCfg := NewBaseConfig(
		WithBaseConfigIsMonitor(true),
		WithBaseConfigCustomConf(&redis.RedisConfig{}),
		)
	l := make(map[string]string)
	l["test"] = "/Users/liufuhong/github-program/grpc-battle/pkg/config"
	baseCfg.Load(l)
	baseCfg.StartMonitor()
}


