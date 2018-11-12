package group

import (
	"local/biz"
	"local/biz/mdl"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	vld "gopkg.in/go-playground/validator.v9"
)

type repoImpl struct {
	db *pg.DB
}

func (r repoImpl) Create(model *mdl.Group) error {
	v := vld.New()
	vldErr := v.Struct(model)
	if vldErr != nil {
		return biz.NewErr(biz.CodeBadRequest, vldErr.Error())
	}
	err := r.db.Insert(model)
	return err
}
func (r repoImpl) ListAll() (*[]mdl.Group, error) {
	var groups []mdl.Group
	err := r.db.Model(&groups).Select()

	return &groups, err

}
func (r repoImpl) DeleteByID(id string) (orm.Result, error) {
	rs, err := r.db.Model(&mdl.Group{}).
		Where("id = ?", id).
		Delete()
	return rs, err
}

func (r repoImpl) DeleteAll() (orm.Result, error) {
	return r.db.Model(&mdl.Group{}).
		Where("1=1").
		Delete()
}

func (r repoImpl) ListAllWhereUserIn(userID uint32) (*[]mdl.Group, error) {
	sql := `
	select g.* from groups as g
	left join user_groups ug on ug.group_id = g.id
	where ug.user_id = ?
	`
	var groups []mdl.Group
	_, err := r.db.Query(&groups, sql, userID)
	return &groups, err
}
