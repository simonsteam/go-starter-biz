package boot

import (
	"github.com/go-pg/pg"
	"local/biz"
	"local/biz/utl"
)

// some const
const (
	BootConditionCfgDone biz.BootCondition = "cfg_done"
)

var (
	// CfgModule config module ,provide configs
	CfgModule = biz.Module{
		Name:         "env config",
		Introduction: "read configs and provide config",
		Provider: func() ConfigData {
			return readConfigData()
		},
		BootFn: &biz.BootFunc{
			Preconditions: biz.ZeroBootCondition,
			Fn: func(env *biz.Env) error {
				env.ConditionOK(BootConditionCfgDone)
				return nil
			}},
	}

	// DBModule database module,provide pg.DB according to configs
	DBModule = biz.Module{
		Name:         "DB Module",
		Introduction: "provide db connection",
		Provider: func(cfg *ConfigData) (*pg.DB, error) {
			db := pg.Connect(&pg.Options{
				Database: cfg.DBName,
				User:     cfg.DBUser,
				Password: cfg.DBPassword,
				PoolSize: 2,
			})
			db.OnQueryProcessed(utl.FnDBLogger)
			return db, nil
		},
		BootFn: &biz.BootFunc{
			Preconditions: []biz.BootCondition{},
			Fn: func(db *pg.DB) error {
				//TODO check database tables
				return nil
			},
		},
		CloseFn: func(db *pg.DB) error {
			return db.Close()
		},
	}
)

// ConfigData .
type ConfigData struct {
	DBName     string
	DBUser     string
	DBPassword string
}

func readConfigData() ConfigData {
	// TODO Read config from somewhere
	return ConfigData{
		DBName:     "",
		DBUser:     "",
		DBPassword: "pwd",
	}
}
