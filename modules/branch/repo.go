package branch

import (
	"local/biz/mdl"
)

// RepoI branch repository
type RepoI interface {
	Create(*mdl.Branch) (uint32, error)
	SelectByID(uint32) (*mdl.Branch, error)
	Update(*mdl.Branch) error
	SelectAll() (*[]mdl.Branch, error)
	DeleteByID(uint32) error
}
