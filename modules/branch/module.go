package branch

import (
	"local/biz"
	"log"

	"github.com/go-pg/pg"
)

// Module info
var (
	Module = biz.Module{
		Provider: func(db *pg.DB) RepoI {
			return repoImpl{
				db: db,
			}
		},
		Bootstrap: func() {
			log.Println("Module branch bootstraped")
		},
	}
)
