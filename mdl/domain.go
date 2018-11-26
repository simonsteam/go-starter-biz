package mdl

// Domain any branch belong to 0 or 1 domain,each domain is managed by only one user(manager)
type Domain struct {
	tableName struct{} `sql:"domain"`
	Base
	Name      string    `json:"name" validate:"required"`
	MgrUserID int       `json:"mgrUserID"` // domain manager id
	MgrUser   *User     `json:"mgrUser" pg:"fk:mgr_user_id"`
	Branchs   *[]Branch `json:"branchs"`
}
