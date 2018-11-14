package group_test

import (
	"local/biz/modules/user"
	"testing"

	"local/biz"
	"local/biz/mdl"
	"local/biz/modules/group"
	"local/biz/test"

	"github.com/go-pg/pg"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	env := test.CreateEnv(t, test.GetTestDatabaseNameForCaller(), true)
	defer env.Release(t, false)

	env.ProvideTestDB()
	biz.BootstrapModules(env.C, group.Module)

	cases := []struct {
		Name   string
		Before func(group.RepoI)
		Val    mdl.Group
		Err    bool
	}{
		{
			"should be validate error",
			nil,
			mdl.Group{ID: "GRP1"},
			true,
		},
		{
			"should create success",
			nil,
			test.TestDataVldGroups[0],
			false,
		},
	}

	err := env.C.Invoke(func(repo group.RepoI) {
		for _, c := range cases {
			if c.Before != nil {
				c.Before(repo)
			}
			err := repo.Create(&c.Val)
			if c.Err {
				assert.Error(t, err, c.Name)
			} else {
				assert.Nil(t, err, c.Name)
			}
		}
	})

	assert.Nil(t, err)

}

func TestListDelete(t *testing.T) {
	env := test.CreateEnv(t, test.GetTestDatabaseNameForCaller(), true)
	defer env.Release(t, false)

	env.ProvideTestDB()
	biz.BootstrapModules(env.C, group.Module)

	err := env.C.Invoke(func(repo group.RepoI) {
		testDataLen := len(test.TestDataVldGroups)

		for _, g := range test.TestDataVldGroups {
			err := repo.Create(&g)
			assert.Nil(t, err, "valid data should not create error")
		}
		rs, err := repo.ListAll()
		assert.Nil(t, err, "list all should success, always")
		assert.Equal(t, testDataLen, len(*rs))

		repo.DeleteByID(test.TestDataVldGroups[0].ID)
		rs, err = repo.ListAll()
		assert.Equal(t, len(*rs), testDataLen-1, "should subtract 1")

		repo.DeleteAll()
		rs, err = repo.ListAll()
		assert.Equal(t, len(*rs), 0, "should be all deleted")
	})

	assert.Nil(t, err)
}

func TestListAllWhereUserIn(t *testing.T) {
	env := test.CreateEnv(t, test.GetTestDatabaseNameForCaller(), true)
	defer env.Release(t, test.KeepTestDBNo)

	env.ProvideTestDB()
	biz.BootstrapModules(env.C, user.Module, group.Module)

	err := env.C.Invoke(func(userRepo user.RepoI, repo group.RepoI, db *pg.DB) {
		groups := test.TestDataVldGroups
		err := db.Insert(&groups)
		assert.Nil(t, err)

		u := test.TestDataVldUsers[0]
		id, err := userRepo.Create(&u)
		assert.Nil(t, err)
		assert.True(t, id > 0)

		var groupIDs []string
		for _, g := range groups {
			groupIDs = append(groupIDs, g.ID)
		}
		assert.Equal(t, len(groups), len(groupIDs))
		err = userRepo.SetGroups4User(id, &groupIDs)
		assert.Nil(t, err)

		userGroups, err := repo.ListAllWhereUserIn(id)
		assert.Nil(t, err)
		assert.Equal(t, groups, *userGroups)

	})

	assert.Nil(t, err)
}
