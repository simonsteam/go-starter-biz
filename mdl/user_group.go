package mdl

// UserGroup n2n relation
type UserGroup struct {
	UserID    uint32
	GroupID  string
	User User
	Group Group
}
