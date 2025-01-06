package department

import "github.com/google/uuid"

type Department struct {
	Id        uuid.UUID `db:"department_id"`
	Name      string    `db:"name"`
	ManagerId uuid.UUID `db:"manager_id"`
}

type FilterDepartment struct {
	Limit     int       `db:"limit"`
	Offset    int       `db:"offset"`
	Name      string    `db:"name"`
	ManagerId uuid.UUID `db:"manager_id"`
}

type FilterDepartment struct {
	Limit     uint      `db:"limit"`
	Offset    uint      `db:"offset"`
	Name      string    `db:"name"`
	ManagerId uuid.UUID `db:"manager_id"`
}
