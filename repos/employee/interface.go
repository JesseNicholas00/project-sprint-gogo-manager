package employee

import "context"

type EmployeeRepository interface {
	GetEmployeeByFilters(ctx context.Context, filter FilterEmployee) ([]Employee, error)
}
