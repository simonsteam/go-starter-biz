package user

import (
	"context"
	"local/biz/mdl"
)

// SvsI user service
type SvsI interface {
	Register(context.Context, *RegisterUserParam) error
	SetGroups4User(context.Context, *SetGroups4UserParam) error
	AddUser(context.Context, *mdl.User) (id uint32, err error)
	FindByID(context.Context, uint32) (*mdl.User, error)
}

// SetGroups4UserParam .
type SetGroups4UserParam struct {
	GroupIDs *[]string `validate:"required,gt=1"`
	UserID uint32 `validate:"required"`
}

// RegisterUserParam 注册参数
type RegisterUserParam struct {
	Phone    string `validate:"required"`
	Captch   string `validate:"required"`
	Username string `validate:"required"`
	Password string `validate:"required"`
}
