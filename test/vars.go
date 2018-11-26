package test

import (
	"local/biz/ac"
	"local/biz/mdl"
	"local/biz/modules/config"
	"local/biz/modules/demo"
)

// some test datas
const (
	TestDBUser   = "biz_test"
	TestPassword = "pwd"

	DropTestDB    = true
	NotDropTestDB = false
)

// some permissions
var (
	ConfigPermissions   []ac.Permission = []ac.Permission{config.PermissionCreateConfig, config.PermissionReadConfig, config.PermissionEditConfig}
	DemoDataPermissions []ac.Permission = []ac.Permission{demo.PermissionUpdateContent, demo.PermissionReadData}

	// AllPermissions []string = append(ConfigPermissions, DemoDataPermissions...)
	// AllStringPermissions []string = AllPermissions
)

// some test datas, should be used for testing only
var (
	TestDataVldGroups = []mdl.Group{{
		ID:          "TGROUP2",
		Name:        "nm",
		Permissions: []string{"ADMIN", "USER"},
	}, {
		ID:          "TGROUP3",
		Name:        "管理员",
		Permissions: []string{"ADMIN", "USER", "WATCHER"},
	},
	// {
	// 	ID: "SUPER",
	// 	Name: "超级管理员super admin",
	// 	Permissions: AllPermissions,
	// }
	}

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
