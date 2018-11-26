package demo_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"local/biz"
	"local/biz/modules/boot"
	"local/biz/modules/branch"
	"local/biz/modules/demo"
	"local/biz/modules/group"
	"local/biz/modules/user"
	"local/biz/test"
	"local/biz/utl"
	"testing"
)

func TestUpdateContent(t *testing.T) {
	helper := test.NewHelper(t, test.GetTestDatabaseNameForCaller(), test.DropTestDB)
	defer helper.Close(t, test.NotDropTestDB)

	env := biz.NewEnv(helper.CfgModule, boot.DBModule, user.Module, group.Module, branch.Module, demo.Module)
	env.Boot()
	defer env.Close()

	asrt := assert.New(t)
	err := env.Container.Invoke(func(svs demo.SvsI, userRepo user.RepoI, groupRepo group.RepoI, branchRepo branch.RepoI) {
		//
		// ctx :=

	})
	assert.Nil(t, err, utl.FnErrorString(err))

}
