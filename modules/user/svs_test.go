package user_test

import (
	"github.com/stretchr/testify/assert"
	"local/biz"
	"local/biz/mdl"
	"local/biz/modules/boot"
	"local/biz/modules/group"
	"local/biz/modules/user"
	"local/biz/test"
	"testing"
)

func addUserAndAssert(t *testing.T, svs user.SvsI, u *mdl.User) uint32 {
	id, err := svs.AddUser(nil, u)
	assert.Nil(t, err)
	assert.True(t, id > 0, "id should be gt 0")
	return id
}

func TestRegister(t *testing.T) {
	t.SkipNow()
	t.Error("TBD")
}

// 测试为用户设定角色组..
func TestSetGroups4User(t *testing.T) {
	helper := test.NewHelper(t, test.GetTestDatabaseNameForCaller(), test.DropTestDB)
	defer helper.Close(t, test.DropTestDB)

	env := biz.NewEnv(helper.CfgModule, boot.DBModule, group.Module, user.Module)
	env.Boot()
	defer env.Close()

	err := env.Container.Invoke(func(svs user.SvsI, groupSvs group.SvsI, groupRepo group.RepoI) {
		insertU := &test.TestDataVldUsers[0]
		id := addUserAndAssert(t, svs, insertU)

		var groupIDs []string
		for _, g := range test.TestDataVldGroups {
			err := groupSvs.Create(nil, &g)
			assert.Nil(t, err)
			groupIDs = append(groupIDs, g.ID)
		}

		err := svs.SetGroups4User(nil, &user.SetGroups4UserParam{
			UserID:   id,
			GroupIDs: &groupIDs,
		})
		assert.Nil(t, err)

		userGroups, err := groupRepo.ListAllWhereUserIn(id)
		assert.Nil(t, err)
		assert.Equal(t, len(groupIDs), len(*userGroups))
	})
	assert.Nil(t, err)

}

func TestAddUser(t *testing.T) {
	helper := test.NewHelper(t, test.GetTestDatabaseNameForCaller(), test.DropTestDB)
	defer helper.Close(t, test.DropTestDB)

	env := biz.NewEnv(helper.CfgModule, boot.DBModule, group.Module, user.Module)
	env.Boot()
	defer env.Close()

	err := env.Container.Invoke(func(svs user.SvsI) {
		insertU := &test.TestDataVldUsers[0]
		id := addUserAndAssert(t, svs, insertU)

		u, err := svs.FindByID(nil, id)
		assert.Nil(t, err)
		assert.Equal(t, insertU.Username, u.Username)
	})
	assert.Nil(t, err)
}
