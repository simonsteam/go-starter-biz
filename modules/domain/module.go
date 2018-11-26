package domain

import (
	"context"
	"github.com/go-pg/pg"
	"local/biz"
	"local/biz/ac"
	"local/biz/mdl"

	vld "gopkg.in/go-playground/validator.v9"
)

// some const
const (
	PermissionCreate ac.Permission = "create_domain"
	PermissionRead   ac.Permission = "read_domain"
	PermissionUpdate ac.Permission = "update_domain"
	PermissionDelete ac.Permission = "delete_domain"
)

var (
	// Module see introduction
	Module = biz.Module{
		Name:         "domain module",
		Introduction: "provide domain service",
		Provider: func(db *pg.DB, vldt *vld.Validate) (SvsI, error) {
			r := repo{
				db: db,
			}
			svs := svsImpl{
				repo: r,
				vld:  vldt,
			}
			return svs, nil
		},
	}
)

// SvsI domain service interface
type SvsI interface {
	Create(context.Context, *mdl.Domain) (int, error)
	SelectAll(context.Context) ([]mdl.Domain, error)
	Update(context.Context, *mdl.Domain) error
	DeleteByID(context.Context, int) error
}

// some module internal infos
const (
	ResType = "domain_data"
)
