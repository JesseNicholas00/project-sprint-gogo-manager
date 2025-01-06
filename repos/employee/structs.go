package employee

import "github.com/google/uuid"

type Employee struct {
	IdentityNumber   string    `db:"identity_number"`
	Name             string    `db:"name"`
	EmployeeImageUri string    `db:"employee_image_uri"`
	Gender           string    `db:"gender"`
	DepartmentId     uuid.UUID `db:"department_id"`
}
