package branch_test

import (
	"github.com/stretchr/testify/assert"
	"local/biz"
	"local/biz/modules/boot"
	"local/biz/modules/branch"
	"local/biz/modules/group"
	"local/biz/modules/user"
	"local/biz/test"
	"testing"
	// "local/biz/mdl"
)

func assertNilErrAndIDGt0(t *testing.T, id uint32, err error) {
	assert.Nil(t, err)
	assert.True(t, id > 0)
}

func TestCRUD(t *testing.T) {
	helper := test.NewHelper(t, test.GetTestDatabaseNameForCaller(), test.DropTestDB)
	defer helper.Close(t, test.DropTestDB)

	env := biz.NewEnv(helper.CfgModule, boot.DBModule, user.Module, branch.Module, group.Module)
	env.Boot()
	defer env.Close()

	err := env.Container.Invoke(func(uRepo user.RepoI, repo branch.RepoI) {
		u := test.TestDataVldUsers[0]
		id, err := uRepo.Create(&u)
		assertNilErrAndIDGt0(t, id, err)

		b := test.TestDataVldBranchs[0]
		b.MgrUserID = id
		bID, err := repo.Create(&b)
		assertNilErrAndIDGt0(t, bID, err)

		dbUser, err := uRepo.FindByID(id)
		assert.Nil(t, err)

		dbUser.BranchID = bID
		err = uRepo.Update(dbUser)
		assert.Nil(t, err)

	})
	assert.Nil(t, err)
}
