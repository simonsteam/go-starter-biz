package ac_test

import (
	"github.com/stretchr/testify/assert"
	"local/biz"
	"local/biz/ac"
	"testing"
)

func assertCheckResult(t *testing.T, allow bool, err error) {
	asrt := assert.New(t)
	if allow {
		asrt.Nil(err)
	} else {
		asrt.NotNil(err)
		er, ok := err.(biz.Err)
		asrt.True(ok)
		asrt.Equal(biz.CodeForbidden, er.Code)
	}
}

func TestHasPermission(t *testing.T) {
	cases := []struct {
		ac.Sub
		Allow bool
	}{
		{
			Sub:   ac.Sub{Permissions: []string{"create"}},
			Allow: true,
		},
		{
			Sub:   ac.Sub{Permissions: []string{""}},
			Allow: false,
		},
	}

	rule := ac.Do(ac.CREATE).
		To(ac.Res{}).
		Should(ac.HasPermission("create"))

	for _, c := range cases {
		err := rule.Check(c.Sub)
		assertCheckResult(t, c.Allow, err)
	}
}

func TestInResDomains(t *testing.T) {
	cases := []struct {
		ac.Sub
		Res   ac.Res
		Allow bool
	}{
		{
			Sub:   ac.Sub{Domains: []string{"a", "b", "c"}},
			Allow: true,
			Res:   ac.Res{Domains: []string{"a", "b"}},
		},
		{
			Sub:   ac.Sub{Domains: []string{"a", "b"}},
			Allow: false,
			Res:   ac.Res{Domains: []string{"a", "c"}},
		},
		{
			Sub:   ac.Sub{Domains: []string{"c", "d"}},
			Allow: false,
			Res:   ac.Res{Domains: []string{"a", "b"}},
		},
		{
			Sub:   ac.Sub{Domains: []string{}},
			Allow: true,
			Res:   ac.Res{Domains: []string{}},
		},
	}

	for _, c := range cases {
		err := ac.Do("any_action").
			To(c.Res).
			Should(ac.InResDomains).
			Check(c.Sub)
		assertCheckResult(t, c.Allow, err)
	}
}
