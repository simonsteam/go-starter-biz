package group

import (
	"context"
	"github.com/go-pg/pg/orm"
	"local/biz"
	"local/biz/mdl"

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
}

// RepoI group repository interface
type RepoI interface {
	Create(*mdl.Group) error
	ListAll() (*[]mdl.Group, error)
	ListAllWhereUserIn(int) (*[]mdl.Group, error)
	DeleteByID(string) (orm.Result, error)
	DeleteAll() (orm.Result, error)
}

// SvsI group service interface
type SvsI interface {
	Create(context.Context, *mdl.Group) error
	ListAll(context.Context) (*[]mdl.Group, error)
	DeleteByID(context.Context, string) error
}
