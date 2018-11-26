package domain

import (
	"context"
	vld "gopkg.in/go-playground/validator.v9"
	"local/biz"
	"local/biz/ac"
	"local/biz/mdl"
	"local/biz/utl"
)

type svsImpl struct {
	repo repo
	vld  *vld.Validate
}

var ruleCreateDomain = ac.Do(ac.CREATE).
	To(ac.Res{Type: ResType}).
	Should(ac.HasPermission(PermissionCreate))

func (s svsImpl) Create(ctx context.Context, model *mdl.Domain) (int, error) {
	if vldErr := s.vld.Struct(model); vldErr != nil {
		return 0, vldErr
	}
	sub, ok := ac.GetSubFromContext(ctx)
	if !ok {
		return 0, biz.ErrUnauthorized
	}
	if err := ruleCreateDomain.Check(sub); err != nil {
		return 0, biz.ErrForbidden
	}

	id, err := s.repo.Insert(model)
	return id, err
}

var ruleReadDomain = ac.Do(ac.READ).
	To(ac.Res{Type: ResType}).
	Should(ac.HasPermission(PermissionRead))

func (s svsImpl) SelectAll(ctx context.Context) ([]mdl.Domain, error) {
	sub, ok := ac.GetSubFromContext(ctx)
	if !ok {
		return nil, biz.ErrUnauthorized
	}
	if err := ruleCreateDomain.Check(sub); err != nil {
		return nil, biz.ErrForbidden
	}
	rs := s.repo.SelectAll()
	return rs, nil
}

func (s svsImpl) Update(ctx context.Context, model *mdl.Domain) error {
	sub, ok := ac.GetSubFromContext(ctx)
	if !ok {
		return biz.ErrUnauthorized
	}

	oldVal, err := s.repo.SelectByID(model.ID)
	if oldVal == nil || err != nil {
		return biz.ErrNotExists
	}
	acRule := ac.Do(ac.UPDATE).
		To(ac.Res{
			Owner: utl.FnItoaPtr(oldVal.MgrUserID),
		}).
		Should(ac.BeOwnerOrHasPermission(PermissionUpdate))

	if acErr := acRule.Check(sub); acErr != nil {
		return biz.ErrForbidden
	}

	oldVal.Name = model.Name
	updErr := s.repo.Update(oldVal)
	return updErr
}

var ruleDelete = ac.Do(ac.DELETE).
	To(ac.Res{Type: ResType}).
	Should(ac.HasPermission(PermissionDelete))

func (s svsImpl) DeleteByID(ctx context.Context, id int) error {
	sub, subOk := ac.GetSubFromContext(ctx)
	if !subOk {
		return biz.ErrUnauthorized
	}

	if acErr := ruleDelete.Check(sub); acErr != nil {
		return biz.ErrForbidden
	}
	delErr := s.repo.DeleteByID(id)
	return delErr
}
