package mdl

import ()

// User 最基本的账号数据结构，账号一旦创建就无法在功能上删除，可以标记废弃、离职
type User struct {
	tableName struct{} `sql:"user"`
	Base
	BranchID    uint32    `json:"branchId"` // 分部Id，除系统管理角色外的账号必须指定门店
	Branch      *Branch   `json:"branch" pg:"fk:branch_id"`
	Description *string   `json:"description,omitempty"`            // 店长备注
	No          *string   `json:"no,omitempty"`                     // 员工编号，系统内唯一
	Password    string    `json:"password"`                         // 登陆密码，数据库加密存储字段，明文的密码有安全度限制以保护账号。建议的规则：包含至少6位，由字母数字特殊符号组成，至少包含一位大写字母一位数字一位特殊符号
	RealName    string    `json:"realName"`                         // 真实姓名
	State       UserState `json:"state"`                            // 账号状态，禁用状态的账号无法登陆和使用系统功能, 正常/禁用/关闭
	Phone       string    `json:"phone"`                            // phone
	Username    string    `json:"username"`                         // 用户名，系统内唯一，用于登陆系统，由字母数字和下划线组成，通常建议使用手机号，也可以使用门店+名字拼音+唯一标识。
	Groups      []Group   `json:"groups" pg:"many2many:user_group"` // 职位/角色 列表
}

// UserState 账号状态，禁用状态的账号无法登陆和使用系统功能, 正常/禁用/关闭
type UserState string

// some const
const (
	UserStateForbid UserState = "Forbid"
	UserStateOff    UserState = "Off"
	UserStateOk     UserState = "Ok"
)
