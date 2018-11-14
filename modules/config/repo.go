package config

import (
	"local/biz/mdl"
)

// RepoI .
type RepoI interface {
	Create(*mdl.Config) (uint32, error)
	SelectAll() (*[]mdl.Config, error)
	Update(*mdl.Config) error
}
