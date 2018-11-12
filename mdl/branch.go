package mdl

// Branch 分部
type Branch struct {
	Address          *string     `json:"address,omitempty"`      // 地址
	AdminDescription string      `json:"adminDescription"`       // 管理员备注
	MgrUserID        int64       `json:"branchManagerUserId"`    // 门店店长ID
	ID               int64       `json:"id"`                     // 系统自动生产的ID
	Introduction     *string     `json:"introduction,omitempty"` // 门店简介
	Lat              *string     `json:"lat,omitempty"`          // 门店纬度
	Lng              *string     `json:"lng,omitempty"`          // 门店经度
	Name             string      `json:"name"`                   // 门店名称
	No               string      `json:"no"`                     // 编号，由管理员手工填写
	State            BranchState `json:"state"`                  // 门店状态
	Tel              *string     `json:"tel,omitempty"`          // 门店联系电话
}

// BranchState 门店状态
type BranchState string

// some state
const (
	BranchStateOff BranchState = "BranchStateOff"
	BranchStateOn  BranchState = "BranchStateOn"
)
