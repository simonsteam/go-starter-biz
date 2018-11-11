package biz

import (
	"context"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	// "errors"
	"runtime"
	"strings"
	"time"

	"local/biz/mdl"

	"go.uber.org/dig"
)

const (
	// CtxUserIDKey userID key in context
	CtxUserIDKey = "userIDKey"
)

// GetModDir go.mod directory
func GetModDir() string {
	_, file, _, _ := runtime.Caller(0)
	return file[:strings.LastIndex(file, "/")]
}

// GetUsrFromContext .
func GetUsrFromContext(ctx context.Context) (Usr, bool) {
	usr, ok := ctx.Value(CtxUserIDKey).(Usr)
	return usr, ok
}

// NewErr create a new error with code,msg,time.Now()
func NewErr(code uint32, msg string) Err {
	return Err{
		Code: code,
		Msg:  msg,
		Time: time.Now(),
	}
}

// BootstrapModules .
func BootstrapModules(c *dig.Container, modules...Module) {
	for _, m := range modules {
		c.Provide(m.Provider)
	}
	for _, m := range modules {
		c.Invoke(m.Bootstrap)
	}
}

func MigrationDatabase(db *pg.DB) error {
	opt := &orm.CreateTableOptions{
		IfNotExists: true,
		FKConstraints: true,
	}

	// register m2m relation,注册多对多关系
	// orm.RegisterTable((*mdl.UserGroup)(nil))
	
	for _, m := range []interface{}{
		(*mdl.User)(nil), (*mdl.Group)(nil), (*mdl.UserGroup)(nil),
	} {
		err := db.CreateTable(m, opt)
		if err != nil {
			return err
		}
	}
	return nil
}
