package domain

import (
	"github.com/go-pg/pg"
	"github.com/stretchr/testify/assert"
	"local/biz"
	"local/biz/mdl"
	"local/biz/modules/boot"
	"local/biz/test"
	"local/biz/utl"
	"testing"
)

func TestCRUD(t *testing.T) {
	helper := test.NewHelper(t, "t_domain_repo_crud", test.DropTestDB)
	defer helper.Close(t, test.NotDropTestDB)

	env := biz.NewEnv(helper.CfgModule, boot.DBModule)
	env.Boot()
	defer env.Close()

	asrt := assert.New(t)
	err := env.Container.Invoke(func(db *pg.DB) {
		r := repo{db: db}

		name := "some domain"

		id, err := r.Insert(&mdl.Domain{
			Name: name,
		})
		asrt.Nil(err, "should insert successfully")
		asrt.True(id > 0, "id should be returned")

		d, err := r.SelectByID(id)
		asrt.Nil(err)
		asrt.Equal(name, d.Name)

		name2 := "other domain"
		d.Name = name2
		err = r.Update(d)
		asrt.Nil(err, "should update success")

		d2, err := r.SelectByID(id)
		asrt.Nil(err)
		asrt.Equal(name2, d2.Name)

		r.Insert(&mdl.Domain{Name: "d2"})

		arr := r.SelectAll()
		asrt.Equal(2, len(arr))

		r.DeleteByID(id)

		arr = r.SelectAll()
		asrt.Equal(1, len(arr))

	})

	asrt.Nil(err, utl.FnErrorString(err))
}
