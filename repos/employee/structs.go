package employee

import "github.com/google/uuid"

type Employee struct {
	IdentityNumber   string    `db:"identity_number"`
	Name             string    `db:"name"`
	EmployeeImageUri string    `db:"employee_image_uri"`
	Gender           string    `db:"gender"`
	DepartmentId     uuid.UUID `db:"department_id"`
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
