package domain

import (
	"github.com/go-pg/pg"
	"local/biz/mdl"
)

type repo struct {
	db *pg.DB
}

func (r repo) Insert(model *mdl.Domain) (int, error) {
	_, err := r.db.Model(model).
		Returning("id").
		Insert()
	return model.ID, err
}

func (r repo) SelectByID(id int) (*mdl.Domain, error) {
	model := mdl.Domain{Base: mdl.Base{ID: id}}
	err := r.db.Select(&model)
	return &model, err
}

func (r repo) Update(model *mdl.Domain) error {
	_, err := r.db.Model(model).
		WherePK().
		Update()
	return err
}

func (r repo) SelectAll() []mdl.Domain {
	var domains []mdl.Domain
	r.db.Model(&domains).
		Select()
	return domains
}

func (r repo) DeleteByID(id int) error {
	_, err := r.db.Model(&mdl.Domain{}).
		Where("id = ?", id).
		Delete()
	return err
}
