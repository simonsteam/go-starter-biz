package group

import (
	"local/biz"

	"github.com/go-pg/pg"
)

// Module injection provider and bootstrap
var Module = biz.Module{
	Provider: func(db *pg.DB) (RepoI, SvsI, error) {
		var repo RepoI = repoImpl{
			db: db,
		}
		var svs SvsI = svsImpl{
			repo: repo,
		}
		return repo, svs, nil
	},
	Bootstrap: func(db *pg.DB) {
		// db.CreateTable(&mdl.Group{}, &orm.CreateTableOptions{
		// 	IfNotExists: true,
		// })
	},
}
