package user_test

import (
	"local/biz/modules/group"
	"local/biz"
	"local/biz/mdl"
	"github.com/stretchr/testify/assert"
	"local/biz/test"
	"local/biz/modules/user"
	"testing"
)

func addUserAndAssert(t *testing.T, svs user.SvsI, u *mdl.User) uint32{
	id, err := svs.AddUser(nil, u)
	assert.Nil(t, err)
	assert.True(t, id >0, "id should be gt 0")
	return id
}

func TestRegister(t *testing.T) {
	t.Error("TBD")
}

func TestSetGroups4User(t *testing.T) {
	env := test.CreateEnv(t, test.GetTestDatabaseNameForCaller(), true)
	defer env.Release(t, true)

	env.ProvideTestDB()
	biz.BootstrapModules(env.C, group.Module, user.Module)

	env.C.Invoke(func(svs user.SvsI, groupSvs group.SvsI){
		insertU := &test.TestDataVldUsers[0]
		id := addUserAndAssert(t, svs, insertU)

		var groupIDs []string
		for _, g := range test.TestDataVldGroups {
			err := groupSvs.Create(nil, &g)
			assert.Nil(t, err)
			groupIDs = append(groupIDs, g.ID)
		}

		err := svs.SetGroups4User(nil, &user.SetGroups4UserParam{
			UserID: id,
			GroupIDs: &groupIDs,
		})
		assert.Nil(t, err)
	})

}

func TestAddUser(t *testing.T) {
	env := test.CreateEnv(t, test.GetTestDatabaseNameForCaller(), true)
	defer env.Release(t, true)

	env.ProvideTestDB()
	biz.BootstrapModules(env.C, group.Module, user.Module)
	
	env.C.Invoke(func(svs user.SvsI){
		insertU := &test.TestDataVldUsers[0]
		id := addUserAndAssert(t, svs, insertU)

		u, err := svs.FindByID(nil, id)
		assert.Nil(t, err)
		assert.Equal(t, insertU.Username, u.Username)
	})
}