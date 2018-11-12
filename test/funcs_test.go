package test_test

import (
	"github.com/go-pg/pg"
	"strings"
	"testing"

	"local/biz"
	"local/biz/test"

	"github.com/stretchr/testify/assert"
)

func TestGetTestDatabaseNameForCaller(t *testing.T) {
	testDbName := test.GetTestDatabaseNameForCaller()
	shouldBe := biz.TestDatabasePrefix + "test_funcs_test_" + strings.ToLower("TestGetTestDatabaseNameForCaller")
	assert.Equal(t, testDbName, shouldBe, "Should be equal")
}

func TestCreateEnv(t *testing.T) {
	// t.SkipNow()
	env := test.CreateEnv(t, "tmpdb", true)
	defer env.Release(t, false)

	biz.MigrationDatabase(env.TestDB)

	var i1, i2 int
	env.TestDB.QueryOne(pg.Scan(&i1, &i2), `select ?,?`, 233, 666)
	assert.Equal(t, i1, 233)
	assert.Equal(t, i2, 666)
}

func TestMigrationDatabaseFromSQL(t *testing.T) {
	env := test.CreateEnv(t, "tmpdb", true)
	defer env.Release(t, true)

	err := biz.MigrationDatabaseFromSQL(env.TestDB)
	assert.Nil(t, err)
}
