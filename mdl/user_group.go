package mdl

// UserGroup n2n relation
type UserGroup struct {
	tableName struct{} `sql:"user_group"`
	UserID    int
	GroupID   string
	User      User
	Group     Group
}
