package test

import (
	"fmt"
	"local/biz/utl"
	"runtime"
	"strings"
	"testing"

	"local/biz"
	"local/biz/modules/boot"

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

func fnDropTestDB(t *testing.T, superDB *pg.DB, dropDBName string) {
	sql := fmt.Sprintf("drop database if exists %s", dropDBName)
	_, err := superDB.Exec(sql)
	assert.Nil(t, err)
}
func NewHelper(t *testing.T, testDBName string, dropTestDBFirst bool) Helper {
	superDB := pg.Connect(&pg.Options{
		Database: "biz_test_template",
		User:     TestDBUser,
		Password: TestPassword,
		PoolSize: 2,
	})
	assert.NotNil(t, superDB)

	helper := Helper{
		SuperDB:    superDB,
		TestDBName: testDBName,
	}
	helper.SuperDB.OnQueryProcessed(utl.FnDBLogger)

	if dropTestDBFirst {
		fnDropTestDB(t, helper.SuperDB, helper.TestDBName)
	}

	sql := fmt.Sprintf("create database %s owner %s ", testDBName, TestDBUser)
	_, err := helper.SuperDB.Exec(sql)
	assert.Nil(t, err)

	testDB := pg.Connect(&pg.Options{
		Database: testDBName,
		User:     TestDBUser,
		Password: TestPassword,
		PoolSize: 2,
	})
	biz.MigrationDatabaseFromSQL(testDB) //TODO production env should not migrate database sql
	testDB.Close()

	helper.CfgModule = biz.Module{
		Name:         "TestEnvConfigModule",
		Introduction: "injected as test env config",
		Provider: func() *boot.ConfigData {
			return &boot.ConfigData{
				DBName:     helper.TestDBName,
				DBUser:     TestDBUser,
				DBPassword: TestPassword,
			}
		},
	}
	return helper
}

// Helper helper in testing,create test database and provide test config
type Helper struct {
	SuperDB    *pg.DB
	TestDBName string
	CfgModule  biz.Module
}

// Close close helper,and drop test database
func (helper Helper) Close(t *testing.T, dropTestDB bool) {
	if dropTestDB {
		fnDropTestDB(t, helper.SuperDB, helper.TestDBName)
	}
	err := helper.SuperDB.Close()
	assert.Nil(t, err)
}

// TestEnv .
// Some time,test case needs a standalone env for test,it ownn standalone database.This help test cases run parallelly.
// in the beginning, im considering using docker test, when needed it starts a new database container,but it increasing learning cost and complexity，and may slow down the speed.
//
// lang:zh_CN 有时候会希望单个测试用例运行在被隔离的环境中，有自己独占的数据库，这样方便并行测试.
// 起初也考虑用docker, 必要的时候运行一个新的数据库容器，但会增加学习成本和复杂度，另外速度上也不太理想。
// type TestEnv struct {
// 	biz.Env
// 	ConnDB     *pg.DB // used to create temp test database
// 	TestDBName string
// }

// // Close drop test db (if not keep it),close db connection
// func (env *TestEnv) Close(t *testing.T, keepTestDB bool) {
// 	errs := env.Env.Close()
// 	assert.Equal(t, 0, len(errs))

// 	if !keepTestDB {
// 		env.DropTestDB(t)
// 	}

// 	err := env.ConnDB.Close()
// 	assert.Nil(t, err)
// }

// // DropTestDB .
// func (env *TestEnv) DropTestDB(t *testing.T) {
// 	sql := fmt.Sprintf("drop database if exists %s", env.TestDBName)
// 	_, err := env.ConnDB.Exec(sql)
// 	assert.Nil(t, err)
// }

// // CreateEnv create new temp database for test
// func CreateEnv(t *testing.T, testDBName string, dropTestDBFirst bool) *TestEnv {
// 	connDB := pg.Connect(&pg.Options{
// 		Database: "biz_test_template",
// 		User:     TestDBUser,
// 		Password: TestPassword,
// 		PoolSize: 2,
// 	})
// 	assert.NotNil(t, connDB)

// 	env := Env{
// 		ConnDB:     connDB,
// 		TestDBName: testDBName,
// 		C:          dig.New(),
// 	}

// 	env.ConnDB.OnQueryProcessed(utl.FnDBLogger)

// 	if dropTestDBFirst {
// 		env.DropTestDB(t)
// 	}

// 	sql := fmt.Sprintf("create database %s owner %s ", testDBName, TestDBUser)
// 	_, err := connDB.Exec(sql)
// 	assert.Nil(t, err)

// 	env.TestDB = pg.Connect(&pg.Options{
// 		Database: testDBName,
// 		User:     TestDBUser,
// 		Password: TestPassword,
// 	})
// 	assert.NotNil(t, env.TestDB)

// 	env.TestDB.OnQueryProcessed(utl.FnDBLogger)

// 	err = biz.MigrationDatabaseFromSQL(env.TestDB)
// 	assert.Nil(t, err)

// 	return &env
// }
