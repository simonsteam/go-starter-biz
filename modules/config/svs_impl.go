package config

import (
	"context"
	"local/biz"
	"local/biz/mdl"
)

type serviceImpl struct {
	repo RepoI
}

var ruleCreateConfig = biz.Do{biz.CREATE}.
	To(biz.Res{}).
	Should(biz.HasPermission(PermissionCreateConfig))

func (svs serviceImpl) Create(ctx context.Context, cfg *mdl.Config) (uint32, error) {
	sub, ok := biz.GetSubFromContext(ctx)
	if !ok {
		return 0, biz.ErrUnauthorized
	}

	if err := ruleCreateConfig.Check(sub); err != nil {
		return 0, err
	}

	// TODO validate
	return svs.repo.Create(cfg)
}

var ruleReadConfig = biz.Do{biz.READ}.
	To(biz.Res{Type: "config"}).
	Should(biz.HasPermission(PermissionReadConfig))

func (svs serviceImpl) SelectAll(ctx context.Context) (*[]mdl.Config, error) {
	sub, ok := biz.GetSubFromContext(ctx)
	if !ok {
		return nil, biz.ErrUnauthorized
	}

	if err := ruleReadConfig.Check(sub); err != nil {
		return nil, err
	}

	return svs.repo.SelectAll()
}

func (svs serviceImpl) Update(ctx context.Context, model *mdl.Config) error {
	sub, ok := biz.GetSubFromContext(ctx)
	if !ok {
		return biz.ErrUnauthorized
	}

	rule := biz.Do{biz.UPDATE}.To(biz.Res{
		ID:   &model.ID,
		Type: ModelType,
	}).Should(biz.HasPermission(PermissionEditConfig))

	if err := rule.Check(sub); err != nil {
		return err
	}

	return svs.repo.Update(model)
}
