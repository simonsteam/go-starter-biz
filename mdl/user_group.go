package mdl

// UserGroup n2n relation
type UserGroup struct {
	tableName struct{} `sql:"user_group"`
	UserID    uint32
	GroupID   string
	User      User
	Group     Group
}
