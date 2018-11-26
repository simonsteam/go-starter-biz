package config

import (
	"local/biz/ac"
	"local/biz/utl"
	// "fmt"
	// "strconv"
	"context"
	"local/biz"
	"local/biz/mdl"
)

type serviceImpl struct {
	repo RepoI
}

var ruleCreateConfig = ac.Do(ac.CREATE).
	To(ac.Res{}).
	Should(ac.HasPermission(PermissionCreateConfig))

func (svs serviceImpl) Create(ctx context.Context, cfg *mdl.Config) (int, error) {
	sub, ok := ac.GetSubFromContext(ctx)
	if !ok {
		return 0, biz.ErrUnauthorized
	}

	if err := ruleCreateConfig.Check(sub); err != nil {
		return 0, err
	}

	// TODO validate
	return svs.repo.Create(cfg)
}

var ruleReadConfig = ac.Do(ac.READ).
	To(ac.Res{Type: "config"}).
	Should(ac.HasPermission(PermissionReadConfig))

func (svs serviceImpl) SelectAll(ctx context.Context) (*[]mdl.Config, error) {
	sub, ok := ac.GetSubFromContext(ctx)
	if !ok {
		return nil, biz.ErrUnauthorized
	}

	if err := ruleReadConfig.Check(sub); err != nil {
		return nil, err
	}

	return svs.repo.SelectAll()
}

func (svs serviceImpl) Update(ctx context.Context, model *mdl.Config) error {
	sub, ok := ac.GetSubFromContext(ctx)
	if !ok {
		return biz.ErrUnauthorized
	}

	rule := ac.Do(ac.UPDATE).
		To(ac.Res{
			ID:   utl.FnItoaPtr(model.ID),
			Type: ModelType,
		}).
		Should(ac.HasPermission(PermissionEditConfig))

	if err := rule.Check(sub); err != nil {
		return err
	}

	return svs.repo.Update(model)
}
