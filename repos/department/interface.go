package department

import "context"

type Repository interface {
	AddDepartment(ctx context.Context, department Department) (Department, error)
}
