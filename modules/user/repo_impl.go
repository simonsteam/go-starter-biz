package user

import (
	// "local/biz"
	"local/biz/mdl"

	"github.com/go-pg/pg"
	// "github.com/go-pg/pg/orm"
	vld "gopkg.in/go-playground/validator.v9"
)

type repoImpl struct {
	db *pg.DB
}

func (r repoImpl) Create(model *mdl.User) (uint32, error) {
	v := vld.New()
	vldErr := v.Struct(model)
	if vldErr != nil {
		return 0, vldErr
	}
	err := r.db.Insert(model)
	return model.ID, err
}

func (r repoImpl) Update(model *mdl.User) error {
	return r.db.Update(model)
}

func (r repoImpl) FindByUsername(username string) (*mdl.User, error) {
	user := new(mdl.User)
	err := r.db.Model(user).
		Where("username = ?", username).
		Select()
	return user, err
}

func (r repoImpl) FindByID(id uint32) (*mdl.User, error) {
	u := &mdl.User{Base: mdl.Base{ID: id}}
	err := r.db.Select(u)
	return u, err
}

func (r repoImpl) SetGroups4User(userID uint32, groupIDs *[]string) error {
	var userGroups []mdl.UserGroup
	for _, gid := range *groupIDs {
		userGroups = append(userGroups, mdl.UserGroup{
			UserID:  userID,
			GroupID: gid,
		})
	}
	return r.db.RunInTransaction(func(tx *pg.Tx) error {
		_, err := r.db.Model(&mdl.UserGroup{}).
			Where("user_id = ?", userID).
			Delete()
		if err != nil {
			return err
		}
		return r.db.Insert(&userGroups)
	})
}
