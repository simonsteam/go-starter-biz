package demo

import (
	"context"
	"github.com/go-pg/pg"
	"local/biz"
	"local/biz/ac"
	"local/biz/mdl"
	"local/biz/utl"
)

type svsImpl struct {
	db *pg.DB
}

func (s svsImpl) SelectByID(ctx context.Context, id int) (*mdl.DemoData, error) {
	model := new(mdl.DemoData)
	err := s.db.Model(&model).
		Where("id = ?", id).
		Select()
	return model, err
}

func (s svsImpl) UpdateContent(ctx context.Context, id int, content string) error {
	sub, subOK := ac.GetSubFromContext(ctx)
	if !subOK {
		return biz.ErrUnauthorized
	}
	data := mdl.DemoData{Base: mdl.Base{ID: id}}
	err := s.db.Select(&data)

	if err != nil {
		return biz.NewErr(400, "Could not find DemoData")
	}

	err = ac.Do(ac.UPDATE).
		To(ac.Res{Owner: utl.FnItoaPtr(data.OwnerID)}).
		Should(ac.BeOwnerOrHasPermission(PermissionUpdateContent)).
		Check(sub)

	if err != nil {
		return biz.ErrForbidden
	}

	_, err = s.db.Model(&data).
		Set("content = ?", content).
		Where("id = ?", id).
		Update()
	if err != nil {
		return biz.NewErr(biz.CodeInternalError, "could not update")
	}
	return nil
}

func (s svsImpl) SelectByBranchIDs(ctx context.Context, branchIDs []int) (*[]mdl.DemoData, error) {
	sub, subOK := ac.GetSubFromContext(ctx)
	if !subOK {
		return nil, biz.ErrUnauthorized
	}

	rule := ac.Do(ac.READ).
		To(ac.Res{Domains: utl.IntsToStrings(branchIDs)}).
		Should(ac.InResDomains)
	if err := rule.Check(sub); err != nil {
		return nil, err
	}

	var datas []mdl.DemoData
	err := s.db.Model(&datas).
		Where("branch_id in (?)", pg.In(branchIDs)).
		Select()

	return &datas, err
}
