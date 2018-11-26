package user

import (
	// "local/biz"
	"local/biz/mdl"

	"github.com/go-pg/pg"
	// "github.com/go-pg/pg/orm"
)

type repoImpl struct {
	db *pg.DB
}

func (r repoImpl) Create(model *mdl.User) (int, error) {
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

func (r repoImpl) FindByID(id int) (*mdl.User, error) {
	u := &mdl.User{Base: mdl.Base{ID: id}}
	err := r.db.Select(u)
	return u, err
}

func (r repoImpl) SetGroups4User(userID int, groupIDs *[]string) error {
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
