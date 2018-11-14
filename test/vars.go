package test

import (
	"local/biz/mdl"
)

// some test datas
const (
	TestDBUser   = "biz_test"
	TestPassword = "pwd"

	DropTestDBBeforeStart     = true
	DontDropTestDBBeforeStart = false
	KeepTestDBYes             = true
	KeepTestDBNo              = false
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

	TestDataVldConfigs = []mdl.Config{
		{
			Key:    "testCfg",
			Type:   mdl.ConfigTypeString,
			Name:   "测试name",
			Module: "test",
		},
	}

	TestDataVldBranchs = []mdl.Branch{
		{
			No:        "BJ001",
			Name:      "001Name",
			Address:   "xxx Street",
			Tel:       "101-001-002",
			AdminDesc: "first branch",
			State:     mdl.BranchStateOff,
		},
		{
			No:        "BJ002",
			Name:      "002Name",
			Address:   "xxx Street2",
			Tel:       "101-001-003",
			AdminDesc: "second branch",
			State:     mdl.BranchStateOff,
		},
	}
)
