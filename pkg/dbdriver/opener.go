package dbdriver

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go/log"
	"grpc-battle/pkg/config"
)

type DRIVERTYPE int

const (
	MYSQL DRIVERTYPE = iota
	// TODO and so on
)

type Opener struct {
	conf   *config.BaseConfig
	name   string
	source string
}

func NewOpener(serverName string) *Opener {
	svrCfg := config.NewBaseConfig(config.WithBaseConfigIsMonitor(false))
	svrCfg.Load(map[string]string{
		serverName: ".",
		serverName: "$HOME/." + serverName,
		serverName: "/etc",
		serverName: "$HOME/etc",
	})
	cs := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=%s&parseTime=true&loc=Local",
		svrCfg.GetString("userName"), svrCfg.GetString("Password"), svrCfg.GetString("Protocol"), svrCfg.GetString("Host"),
		svrCfg.GetString("Port"), svrCfg.GetString("Database"), svrCfg.GetString("Charset"))
	return &Opener{conf: svrCfg, name: serverName, source: cs}
}

func (o *Opener) binder() string {
	return fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=%s&parseTime=true&loc=Local",
		o.conf.GetString("userName"), o.conf.GetString("Password"), o.conf.GetString("Protocol"), o.conf.GetString("Host"),
		o.conf.GetString("Port"), o.conf.GetString("Database"), o.conf.GetString("Charset"))
}

func (o *Opener) OpenDb(driverType DRIVERTYPE) (*sqlx.DB, error) {

	switch driverType {
	case MYSQL:
		db, err := sqlx.Open(o.name, o.source)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		return db, nil
	}
}

// TODO handle the query list
