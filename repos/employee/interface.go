package employee

import "context"

type EmployeeRepository interface {
	AddEmployee(ctx context.Context, employee Employee) (Employee, error)
}
