package employee

import (
	"context"

	"github.com/google/uuid"
)

type EmployeeRepository interface {
	GetEmployeeByFilters(ctx context.Context, filter FilterEmployee) ([]Employee, error)
	AddEmployee(ctx context.Context, employee Employee, userId uuid.UUID) error
	UpdateEmployee(ctx context.Context, employee Employee, identityNumber, userID string) (Employee, error)
	FindEmployeeByIdentityNumber(ctx context.Context, identityNumber, userID string) (Employee, error)
	IsIdentityNumberExist(ctx context.Context, identityNumber, userID string) (bool, error)
}
