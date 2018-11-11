package group_test

import (
	"testing"

	"local/biz"
	"local/biz/mdl"
	"local/biz/modules/group"
	"local/biz/test"

	// "github.com/go-pg/pg"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	env := test.CreateEnv(t, test.GetTestDatabaseNameForCaller(), true)
	defer env.Release(t, false)

	env.ProvideTestDB()
	biz.BootstrapModules(env.C, group.Module)

	cases := []struct{
		Name string
		Before func(group.RepoI)
		Val mdl.Group
		Err bool
	}{
		{
			"should be validate error", 
			nil, 
			mdl.Group{ ID: "GRP1", },
			true,
		},
		{
			"should create success",
			nil,
			test.TestDataVldGroups[0],
			false,
		},
	}

	env.C.Invoke(func(repo group.RepoI) {
		for _, c := range cases {
			if c.Before != nil {
				c.Before(repo)
			}
			err := repo.Create(&c.Val)
			if c.Err {
				assert.Error(t, err, c.Name)
			}else {
				assert.Nil(t, err, c.Name)
			}
		}
	})

}

func TestListDelete(t *testing.T) {
	env := test.CreateEnv(t, test.GetTestDatabaseNameForCaller(), true)
	defer env.Release(t, false)

	env.ProvideTestDB()
	biz.BootstrapModules(env.C, group.Module)

	env.C.Invoke(func(repo group.RepoI) {
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
		assert.Equal(t, len(*rs), testDataLen - 1, "should subtract 1")

		repo.DeleteAll()
		rs, err = repo.ListAll()
		assert.Equal(t, len(*rs), 0, "should be all deleted")
	})
}
