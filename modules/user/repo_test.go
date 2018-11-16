package user_test

import (
	"local/biz/modules/boot"
	"testing"

	"local/biz"
	"local/biz/test"
	// "local/biz/mdl"
	"local/biz/modules/group"
	"local/biz/modules/user"

	// "github.com/go-pg/pg"
	"github.com/stretchr/testify/assert"
)

func TestRepoCreate(t *testing.T) {
	helper := test.NewHelper(t, test.GetTestDatabaseNameForCaller(), test.DropTestDB)
	defer helper.Close(t, test.DropTestDB)

	env := biz.NewEnv(helper.CfgModule, boot.DBModule, group.Module, user.Module)
	env.Boot()
	defer env.Close()

	err := env.Container.Invoke(func(repo user.RepoI, groupRepo group.RepoI) {

		for _, g := range test.TestDataVldGroups {
			err := groupRepo.Create(&g)
			assert.Nil(t, err)
		}

		for _, u := range test.TestDataVldUsers {
			id, err := repo.Create(&u)
			assert.Nil(t, err)
			assert.False(t, id == 0)

			user, err := repo.FindByID(id)
			assert.Nil(t, err)
			assert.Equal(t, u.Username, user.Username)

			// assert.Equal(t, len(u.Groups), len(user.Groups),)
		}

	})
	assert.Nil(t, err)
}
