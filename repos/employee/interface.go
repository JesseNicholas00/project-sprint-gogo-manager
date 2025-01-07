package employee

import "context"

type EmployeeRepository interface {
	GetEmployeeByFilters(ctx context.Context, filter FilterEmployee) ([]Employee, error)
	AddEmployee(ctx context.Context, employee Employee) (Employee, error)
	UpdateEmployee(ctx context.Context, employee Employee, identityNumber string) (Employee, error)
	FindEmployeeByIdentityNumber(ctx context.Context, identityNumber string) (Employee, error)
}
