package group

import (
	"context"
	
	"local/biz/mdl"
)

// SvsI group service interface
type SvsI interface {
	Create(context.Context, *mdl.Group) error
	ListAll(context.Context) (*[]mdl.Group, error)
	DeleteByID(context.Context, string) error
}
