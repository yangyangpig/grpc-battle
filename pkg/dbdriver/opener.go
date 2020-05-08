package dbdriver

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"grpc-battle/pkg/config"
	"log"
)

type DRIVERTYPE int

const (
	MYSQL DRIVERTYPE = iota
	// TODO and so on
)

var (
	ERR_DRIVER = errors.New("driver type error")
)

type Opener struct {
	Conf   *config.BaseConfig
	Name   string
	Source string
	Driver interface{}
}

func NewOpener(serverName string, driverType DRIVERTYPE) (*Opener, error) {
	svrCfg := config.NewBaseConfig(config.WithBaseConfigIsMonitor(false))
	svrCfg.Load(map[string]string{
		serverName: ".",
		serverName: "$HOME/." + serverName,
		serverName: "/etc",
		serverName: "$HOME/etc",
	})
	cs := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		svrCfg.GetString("userName"), svrCfg.GetString("Password"), svrCfg.GetString("Protocol"), svrCfg.GetString("Host"),
		svrCfg.GetString("Port"), svrCfg.GetString("Database"), svrCfg.GetString("Charset"))

	var db interface{}
	var err error
	switch driverType {
	case MYSQL:
		db, err = sqlx.Open(serverName, cs)
		if err != nil {
			log.Fatalf("open %d fail : %+v", serverName, err)
			return nil, err
		}
		// make sure the db is available and put conn into the pool
		if sqx, ok := db.(*sqlx.DB); ok {
			// you can establishing more than one conn and put them into pool
			err = sqx.Ping()
			if err != nil {
				log.Fatalf("ping %d db fail : %+v", serverName, err)
				return nil, err
			}
		} else {
			return nil, ERR_DRIVER
		}

	}

	return &Opener{Conf: svrCfg, Name: serverName, Source: cs, Driver: db}, nil
}
