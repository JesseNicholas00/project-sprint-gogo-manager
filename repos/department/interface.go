package department

import "context"

type DepartmentRepository interface {
	AddDepartment(ctx context.Context, department Department) error
	GetDepartment(ctx context.Context, filter FilterDepartment) ([]Department, error)
}
