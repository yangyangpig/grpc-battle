package redis

import (
	"fmt"
	"time"
)

var (
	defaultConfig = RedisConfig{
		Host:         "127.0.0.1",
		Port:         6379,
		Protocol:     "tcp",
		PingInterval: time.Minute,
		PoolSize:     10,
	}
)

type RedisConfig struct {
	Host         string        `json:"host" yaml:"host"`
	Port         int           `json:"port" yaml:"port"`
	Protocol     string        `json:"protocol" yaml:"protocol"`
	PingInterval time.Duration `json:"ping_interval" yaml:"pingInterval"`
	PoolSize     int           `json:"pool_size" yaml:"poolSize"`
	Password     string        `json:"password" yaml:"password"`
	Database     int           `json:"database" yaml:"database"`
}

func Default() RedisConfig {
	return defaultConfig
}

func (cfg *RedisConfig) Addr() string {
	return fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
}

func (cfg *RedisConfig) IdleCheckFrequency() time.Duration {
	return time.Duration(cfg.PingInterval) * time.Second
}

func (cfg *RedisConfig) Print() string {
	return fmt.Sprintf("host(%v) port(%v) ping_interval(%d) pool_size(%d) protocol(%v) password(%v) database(%d)",
		cfg.Host, cfg.Port, cfg.PingInterval, cfg.PoolSize, cfg.Protocol, cfg.Password, cfg.Database)
}

func (cfg *RedisConfig) UpdateFun() {
	fmt.Println(cfg.Print())
}
