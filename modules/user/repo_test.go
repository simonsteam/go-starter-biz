package user_test

import (
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
	env := test.CreateEnv(t, test.GetTestDatabaseNameForCaller(), true)
	defer env.Release(t, false)

	env.ProvideTestDB()
	biz.BootstrapModules(env.C, group.Module, user.Module)

	env.C.Invoke(func(repo user.RepoI, groupRepo group.RepoI) {

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

}
