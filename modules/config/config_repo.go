package config

import (
	"github.com/go-pg/pg"
	"local/biz/mdl"
)

type repoImpl struct {
	db *pg.DB
}

func (r repoImpl) Create(model *mdl.Config) (int, error) {
	_, err := r.db.Model(model).
		Returning("id").
		Insert()
	return model.Base.ID, err
}

func (r repoImpl) SelectAll() (*[]mdl.Config, error) {
	var configs []mdl.Config
	err := r.db.Model(&configs).Select()
	return &configs, err
}

func (r repoImpl) Update(model *mdl.Config) error {
	_, err := r.db.Model(model).
		WherePK().
		Update()
	return err
}
