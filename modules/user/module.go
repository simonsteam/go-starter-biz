package user

import (
	"local/biz"
	"local/biz/modules/group"
	// "local/biz/mdl"

	"github.com/go-pg/pg"
	// "github.com/go-pg/pg/orm"
	vld "gopkg.in/go-playground/validator.v9"
)

// Module injection provider and bootstrap
var Module = biz.Module{
	Provider: func(db *pg.DB, groupRepo group.RepoI) (RepoI, SvsI, error) {
		var impl RepoI = repoImpl{
			db: db,
		}
		var svs SvsI = svsImpl{
			repo:      impl,
			groupRepo: groupRepo,
			vld:       *vld.New(), // add custom rule when need
		}
		return impl, svs, nil
	},
	Bootstrap: func(db *pg.DB) {
		// db.CreateTable(&mdl.User{}, &orm.CreateTableOptions{
		// 	IfNotExists: true,
		// 	// FKConstraints: true,
		// })
	},
}
