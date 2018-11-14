package mdl

// Branch 分部
type Branch struct {
	tableName struct{} `sql:"branch"`
	Base
	No           string      `json:"no" sql:",notnull"`   // 编号，由管理员手工填写
	Name         string      `json:"name" sql:",notnull"` // 门店名称
	MgrUserID    uint32      `json:"mgrUserId"`           // 门店店长ID
	MgrUser      *User       `json:"mgrUser" pg:"fk:mgr_user_id"`
	Address      string      `json:"address,omitempty" sql:",notnull"`      // 地址
	Tel          string      `json:"tel,omitempty" sql:",notnull"`          // 门店联系电话
	AdminDesc    string      `json:"adminDesc" sql:",notnull"`              // 管理员备注
	Introduction string      `json:"introduction,omitempty" sql:",notnull"` // 门店简介
	State        BranchState `json:"state"`                                 // 门店状态
	Lat          *float64    `json:"lat,omitempty"`                         // 门店纬度
	Lng          *float64    `json:"lng,omitempty"`                         // 门店经度
}

// BranchState 门店状态
type BranchState string

// some state
const (
	BranchStateOff BranchState = "off"
	BranchStateOn  BranchState = "on"
)
