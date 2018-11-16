package test_test

import (
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
