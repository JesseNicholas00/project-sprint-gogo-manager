package employee

import (
	"context"

	"github.com/google/uuid"
)

type EmployeeService interface {
	GetEmployeeByFilters(ctx context.Context, params GetEmployeeReq, res *GetEmployeeResp, userId string) error
	AddEmployee(ctx context.Context, req AddEmployeeReq, res *AddEmployeeRes, userId uuid.UUID) error
	IsIdentityNumberExist(ctx context.Context, identityNumber, userID string) (bool, error)
	UpdateEmployee(ctx context.Context, req UpdateEmployeeReq, res *AddEmployeeRes) error
}
