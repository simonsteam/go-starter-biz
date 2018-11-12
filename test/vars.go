package test

import (
	"local/biz/mdl"
)

// some test datas
const (
	TestDBUser   = "biz_test"
	TestPassword = "pwd"
)

// some test datas, should be used for testing only
var (
	TestDataVldGroups = []mdl.Group{{
		ID:    "TGROUP2",
		Name:  "nm",
		Roles: []string{"ADMIN", "USER"},
	}, {
		ID:    "TGROUP3",
		Name:  "管理员",
		Roles: []string{"ADMIN", "USER", "WATCHER"},
	}}

	TestDataVldUsers = []mdl.User{{
		Username: "userOK1",
		Groups:   TestDataVldGroups,
		Password: "pwd",
		State:    mdl.UserStateOk,
	}}
)
