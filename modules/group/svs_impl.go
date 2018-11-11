package group

import (
	"context"
	
	
	"local/biz/mdl"
)

type svsImpl struct {
	repo RepoI
}

func (svs svsImpl) Create(ctx context.Context, group *mdl.Group) error {
	// TODO acl,log
	err := svs.repo.Create(group)
	return err
}

func (svs svsImpl) ListAll(ctx context.Context) (*[]mdl.Group, error) {
	// TODO acl
	return svs.repo.ListAll()
}

func (svs svsImpl) DeleteByID(ctx context.Context, id string) error {
	// TODO ACL,log
	_, err := svs.repo.DeleteByID(id)
	return err
}
