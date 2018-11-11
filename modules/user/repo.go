package user

import (
	"local/biz/mdl"
	// "github.com/go-pg/pg/orm"
)

// RepoI .
type RepoI interface {
	Create(*mdl.User) (uint32, error)
	Update(*mdl.User) error
	FindByUsername(string) (*mdl.User, error)
	FindByID(uint32) (*mdl.User, error)
	SetGroups4User(userID uint32, groupIDs *[]string) error
}
