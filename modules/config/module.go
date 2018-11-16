package config

import (
	"github.com/go-pg/pg"
	"local/biz"
	"local/biz/mdl"
	// ,"log"
)

var Module = biz.Module{
	Provider: []interface{}{
		func(db *pg.DB) RepoI {
			return repoImpl{
				db: db,
			}
		},
	},
}

// RepoI .
type RepoI interface {
	Create(*mdl.Config) (uint32, error)
	SelectAll() (*[]mdl.Config, error)
	Update(*mdl.Config) error
}
