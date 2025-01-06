package department

import "github.com/google/uuid"

type Department struct {
	Id        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	ManagerId uuid.UUID `db:"manager_id"`
}
