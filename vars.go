package biz

// some common codes
const (
	CodeBadRequest   = 400
	CodeNotExists    = 404
	CodeUnauthorized = 401
	CodeForbidden    = 403

	TestDatabasePrefix = "t_"
)

// RoleName .
type RoleName string

// some roles
const (
	RoleAdmin  RoleName = "admin"
	RoleHuman  RoleName = "human"
	RoleSystem RoleName = "system"
)

// some common errors, should be readonly
var (
	ErrNotExists    = Err{Code: CodeNotExists, Msg: "Target not exists!"}
	ErrUnauthorized = Err{Code: CodeUnauthorized, Msg: "Unauthorized"}
	ErrForbidden    = Err{Code: CodeForbidden, Msg: "You are not permitted to this!"}
)

// some vars
var (
	AllRoleNames      = []RoleName{RoleAdmin, RoleHuman, RoleSystem}
	ZeroBootCondition = []BootCondition{}
)
