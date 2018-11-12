package test

import (
	"fmt"
	"go.uber.org/dig"
	"local/biz/utl"
	"runtime"
	"strings"
	"testing"

	"local/biz"

	"github.com/go-pg/pg"
	"github.com/stretchr/testify/assert"
)

// GetTestDatabaseNameForCaller .
// build test database name for caller test \\ lang:zh_CN 为调用的测试方法拼接一个数据库名字
// eg: you call this method from \\ lang:zh_CN 如果你从这个目录调用这个方法
// ${workspace}/modules/user_test.go#TestCreate
//  then you got \\ lang:zh_CN 那么将返回
// t_modules_user_test_testcreate
func GetTestDatabaseNameForCaller() string {
	pc, file, _, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)

	modDir := biz.GetModDir()

	relPath := string(file[len(modDir) : len(file)-3])
	funcName := f.Name()[strings.LastIndex(f.Name(), "."):]
	res := strings.Trim(relPath+funcName, "/")
	res = strings.Replace(res, "/", "_", -1)
	res = strings.Replace(res, ".", "_", -1)
	return strings.ToLower(biz.TestDatabasePrefix + res)
}

// Env .
// Some time,test case needs a standalone env for test,it ownn standalone database.This help test cases run parallelly.
// in the beginning, im considering using docker test, when needed it starts a new database container,but it increasing learning cost and complexity，and may slow down the speed.
//
// lang:zh_CN 有时候会希望单个测试用例运行在被隔离的环境中，有自己独占的数据库，这样方便并行测试.
// 起初也考虑用docker, 必要的时候运行一个新的数据库容器，但会增加学习成本和复杂度，另外速度上也不太理想。
type Env struct {
	ConnDB     *pg.DB // used to create temp test database
	TestDB     *pg.DB // database actually used in test
	TestDBName string
	C          *dig.Container // dependency injection container
}

// ProvideTestDB provide temp test database as db provider
func (env *Env) ProvideTestDB() {
	env.C.Provide(func() *pg.DB {
		return env.TestDB
	})
}

// Release drop test db (if not keep it),close db connection
func (env *Env) Release(t *testing.T, keepTestDB bool) {
	err := env.TestDB.Close()
	assert.Nil(t, err)

	if !keepTestDB {
		env.DropTestDB(t)
	}

	err = env.ConnDB.Close()
	assert.Nil(t, err)
}

// DropTestDB .
func (env *Env) DropTestDB(t *testing.T) {
	sql := fmt.Sprintf("drop database if exists %s", env.TestDBName)
	_, err := env.ConnDB.Exec(sql)
	assert.Nil(t, err)
}

// CreateEnv create new temp database for test
func CreateEnv(t *testing.T, testDBName string, dropTestDBFirst bool) *Env {
	connDB := pg.Connect(&pg.Options{
		Database: "biz_test_template",
		User:     TestDBUser,
		Password: TestPassword,
		PoolSize: 2,
	})
	assert.NotNil(t, connDB)

	env := Env{
		ConnDB:     connDB,
		TestDBName: testDBName,
		C:          dig.New(),
	}

	env.ConnDB.OnQueryProcessed(utl.FnDBLogger)

	if dropTestDBFirst {
		env.DropTestDB(t)
	}

	sql := fmt.Sprintf("create database %s owner %s ", testDBName, TestDBUser)
	_, err := connDB.Exec(sql)
	assert.Nil(t, err)

	env.TestDB = pg.Connect(&pg.Options{
		Database: testDBName,
		User:     TestDBUser,
		Password: TestPassword,
	})
	assert.NotNil(t, env.TestDB)

	env.TestDB.OnQueryProcessed(utl.FnDBLogger)

	biz.MigrationDatabaseFromSQL(env.TestDB)

	return &env
}
