package demo

import (
	"context"
	"github.com/go-pg/pg"
	"local/biz"
	"local/biz/ac"
	"local/biz/mdl"
)

// Module provide demoData service
var Module = biz.Module{
	Name:         "Demo module",
	Introduction: "Demonstrate using domain based access control,owner access control",
	Provider: func(db *pg.DB) (SvsI, error) {
		// TODO
		return svsImpl{db: db}, nil
	},
}

// some fixed permissions
const (
	PermissionUpdateContent ac.Permission = "update_content"
	PermissionReadData      ac.Permission = "read_data"
)

// SvsI demo data service interface
type SvsI interface {
	SelectByID(context.Context, int) (*mdl.DemoData, error)
	// UpdateContent access control: should be the owner of data or has update permission
	UpdateContent(context.Context, int, string) error
	// SelectByBranchIDs access control: should in domains and has read permission
	SelectByBranchIDs(context.Context, []int) (*[]mdl.DemoData, error)
}
