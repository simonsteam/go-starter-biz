package config

import (
	"context"
	"github.com/go-pg/pg"
	"local/biz"
	"local/biz/mdl"
	// ,"log"
)

// some const
const (
	PermissionCreateConfig = "create_config"
	PermissionReadConfig   = "read_config"
	PermissionEditConfig   = "edit_config"
	ModelType              = "config"
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

// SvsI service interface
type SvsI interface {
	Create(context.Context, *mdl.Config) (uint32, error)
	SelectAll(context.Context) (*[]mdl.Config, error)
	Update(context.Context, *mdl.Config) error
}
