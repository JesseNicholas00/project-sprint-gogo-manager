package role

type Role int

const (
	User = iota
)

func GetRole() Role {
	return User
}
