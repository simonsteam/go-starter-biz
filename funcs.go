package biz

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"errors"
	"io/ioutil"
	"log"
	"runtime"
	"strings"
	"time"

	"local/biz/mdl"
)

// GetModDir go.mod directory
func GetModDir() string {
	_, file, _, _ := runtime.Caller(0)
	return file[:strings.LastIndex(file, "/")]
}

// NewErr create a new error with code,msg,time.Now()
func NewErr(code int, msg string) Err {
	return Err{
		Code: code,
		Msg:  msg,
		Time: time.Now(),
	}
}

func MigrationDatabase(db *pg.DB) error {
	log.Printf("----MigrationDatabase----")
	opt := &orm.CreateTableOptions{
		IfNotExists:   true,
		FKConstraints: true,
	}

	// register m2m relation,注册多对多关系
	// orm.RegisterTable((*mdl.UserGroup)(nil))

	for _, m := range []interface{}{
		(*mdl.User)(nil), (*mdl.Group)(nil), (*mdl.UserGroup)(nil),
	} {
		log.Printf(":create table: %T", m)
		err := db.CreateTable(m, opt)
		if err != nil {
			return err
		}
	}
	return nil
}

// MigrationDatabaseFromSQL .
func MigrationDatabaseFromSQL(db *pg.DB) error {
	content, err := ioutil.ReadFile(GetModDir() + "/db.sql")
	if err != nil {
		return err
	}

	sqls := strings.Split(string(content), "--##")
	for _, sql := range sqls {
		_, e := db.Exec(strings.Trim(sql, "\n"))
		if e != nil {
			return errors.New(sql + e.Error())
		}
	}
	return nil
}
