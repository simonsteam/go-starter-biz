package domain_test

import (
	"github.com/stretchr/testify/assert"
	"local/biz"
	"local/biz/modules/boot"
	"local/biz/modules/domain"
	"local/biz/test"
	"local/biz/utl"
	"testing"
)

func TestCreate(t *testing.T) {
	helper := test.NewHelper(t, test.GetTestDatabaseNameForCaller(), test.DropTestDB)
	defer helper.Close(t, test.NotDropTestDB)

	env := biz.NewEnv(helper.CfgModule, boot.DBModule, boot.ToolModule, domain.Module)
	env.Boot()
	defer env.Close()

	asrt := assert.New(t)
	err := env.Container.Invoke(func(svs domain.SvsI) {
		// t.Error("TBD")
	})
	asrt.Nil(err, utl.FnErrorString(err))

}
