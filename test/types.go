package test

import (
	"github.com/stretchr/testify/suite"
	"local/biz"
	"testing"
)

type OneTestOneDBSuite struct {
	suite.Suite
	TT                 *testing.T
	DBName             string
	DropDBBeforeCreate bool
	DropDBAfterTest    bool
	Modules            []biz.Module
	Helper             Helper
	Env                *biz.Env
}

// SetupSuite create db, boot env
func (s *OneTestOneDBSuite) SetupSuite() {
	s.Helper = NewHelper(s.TT, s.DBName, s.DropDBBeforeCreate)
	s.Env = biz.NewEnv(s.Modules...)
	s.Env.Boot()
}

// TearDownSuite Close env, close helper
func (s *OneTestOneDBSuite) TearDownSuite() {
	s.Env.Close()
	s.Helper.Close(s.TT, s.DropDBAfterTest)
}
