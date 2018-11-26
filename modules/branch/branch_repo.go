package branch

import (
	// "local/biz"
	"local/biz/mdl"

	"github.com/go-pg/pg"
)

type repoImpl struct {
	db *pg.DB
}

func (r repoImpl) Create(model *mdl.Branch) (int, error) {
	_, err := r.db.Model(model).
		Returning("id").
		Insert()
	return model.Base.ID, err
}

func (r repoImpl) SelectByID(id int) (*mdl.Branch, error) {
	model := &mdl.Branch{Base: mdl.Base{ID: id}}
	err := r.db.Select(model)
	return model, err
}

func (r repoImpl) Update(model *mdl.Branch) error {
	_, err := r.db.Model(model).
		WherePK().
		Update()
	return err
}

func (r repoImpl) SelectAll() (*[]mdl.Branch, error) {
	var models []mdl.Branch
	err := r.db.Model(models).Select()
	return &models, err
}

func (r repoImpl) DeleteByID(id int) error {
	_, err := r.db.Model(&mdl.Branch{}).
		Where("id = ?", id).
		Delete()
	return err
}
