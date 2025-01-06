package department

import "context"

type DepartmentRepository interface {
	AddDepartment(ctx context.Context, department Department) error
}
