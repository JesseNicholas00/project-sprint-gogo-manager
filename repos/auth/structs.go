package auth

type User struct {
	Id       string `db:"user_id"`
	Password string `db:"password"`
	Email    string `db:"email"`
}
