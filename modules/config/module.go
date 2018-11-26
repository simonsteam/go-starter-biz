package config

import (
	"context"
	"github.com/go-pg/pg"
	"local/biz"
	"local/biz/ac"
	"local/biz/mdl"
	// ,"log"
)

// some const
const (
	ModelType = "config"
)

// some permissions
const (
	PermissionCreateConfig ac.Permission = "create_config"
	PermissionReadConfig   ac.Permission = "read_config"
	PermissionEditConfig   ac.Permission = "edit_config"
)

// some vars
var (
	Module = biz.Module{
		Provider: []interface{}{
			func(db *pg.DB) RepoI {
				return repoImpl{
					db: db,
				}
			},
		},
	}
)

// RepoI .
type RepoI interface {
	Create(*mdl.Config) (int, error)
	SelectAll() (*[]mdl.Config, error)
	Update(*mdl.Config) error
}

// SvsI service interface
type SvsI interface {
	Create(context.Context, *mdl.Config) (int, error)
	SelectAll(context.Context) (*[]mdl.Config, error)
	Update(context.Context, *mdl.Config) error
}
