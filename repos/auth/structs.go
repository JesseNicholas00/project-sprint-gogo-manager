package auth

type User struct {
	Id              string `db:"user_id"`
	Password        string `db:"password"`
	Email           string `db:"email"`
	Name            string `db:"name"`
	UserImageUri    string `db:"user_image_uri"`
	CompanyName     string `db:"company_name"`
	CompanyImageUri string `db:"company_image_uri"`
}
