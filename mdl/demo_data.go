package mdl

// DemoData demo for Demodata access control
type DemoData struct {
	tableName struct{} `sql:"demo_data"`
	Base
	Branch   Branch `json:"branch" pg:"fk:branch_id"`
	BranchID int    `json:"branchID" ` // that is something like domain
	Owner    User   `json:"user" pg:"fk:owner_id"`
	OwnerID  int    `json:"owner" `
	Content  string `json:"content"`
}
