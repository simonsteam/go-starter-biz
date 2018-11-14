package config

import (
	"github.com/go-pg/pg"
	"local/biz"
	"log"
)

var Module = biz.Module{
	Provider: []interface{}{
		func(db *pg.DB) RepoI {
			return repoImpl{
				db: db,
			}
		},
	},
	Bootstrap: func() {
		log.Println("Module config bootstraped!")
	},
}
