package employee

type Employee struct {
	IdentityNumber   string `db:"identity_number"`
	Name             string `db:"name"`
	EmployeeImageUri string `db:"employee_image_uri"`
	Gender           string `db:"gender"`
	DepartmentId     int    `db:"department_id"`
}

type FilterEmployee struct {
	Limit          int    `db:"limit"`
	Offset         int    `db:"offset"`
	IdentityNumber string `db:"identity_number"`
	Name           string `db:"name"`
	Gender         string `db:"gender"`
	DepartementId  int    `db:"department_id"`
	UserId         string `db:"user_id"`
}
