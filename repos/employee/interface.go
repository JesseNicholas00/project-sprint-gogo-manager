package employee

import (
	"context"

	"github.com/google/uuid"
)

type EmployeeRepository interface {
	GetEmployeeByFilters(ctx context.Context, filter FilterEmployee) ([]Employee, error)
	AddEmployee(ctx context.Context, employee Employee, userId uuid.UUID) error
	UpdateEmployee(ctx context.Context, employee Employee, identityNumber string) (Employee, error)
	FindEmployeeByIdentityNumber(ctx context.Context, identityNumber string) (Employee, error)
	IsIdentityNumberExist(ctx context.Context, identityNumber string) (bool, error)
}
