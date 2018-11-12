package group

import (
	"github.com/go-pg/pg/orm"
	"local/biz/mdl"
)

// RepoI group repository interface
type RepoI interface {
	Create(*mdl.Group) error
	ListAll() (*[]mdl.Group, error)
	ListAllWhereUserIn(uint32) (*[]mdl.Group, error)
	DeleteByID(string) (orm.Result, error)
	DeleteAll() (orm.Result, error)
}
