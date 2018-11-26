package config_test

import (
	"github.com/stretchr/testify/assert"
	"local/biz/modules/boot"
	"testing"

	"local/biz"
	"local/biz/modules/config"
	// "local/biz/mdl"
	"local/biz/test"
)

func TestCRUd(t *testing.T) {
	helper := test.NewHelper(t, "t_config_repo_crud", test.DropTestDB)
	defer helper.Close(t, test.DropTestDB)

	env := biz.NewEnv(helper.CfgModule, boot.DBModule, config.Module)
	env.Boot()
	defer env.Close()

	err := env.Container.Invoke(func(repo config.RepoI) {
		cfg := test.TestDataVldConfigs[0]

		id, err := repo.Create(&cfg)
		assert.Nil(t, err, "should insert successfuly")
		assert.True(t, id > 0, "should generate auto id")

		_, err = repo.Create(&cfg)
		assert.Error(t, err, "should be duplicate key error")

		cfgs, err := repo.SelectAll()
		assert.Nil(t, err, "should select successfy")
		assert.Equal(t, 1, len(*cfgs), "should same as insertd len")

		updCfg := (*cfgs)[0]
		updCfg.Name = "updName"
		assert.Nil(t, repo.Update(&updCfg), "should update succ")

		cfgs, err = repo.SelectAll()
		newCfg := (*cfgs)[0]
		assert.Equal(t, newCfg.Name, updCfg.Name, "new valued should be equals")
	})
	assert.Nil(t, err)
}
