package employee

import "context"

type EmployeeService interface {
	GetEmployeeByFilters(ctx context.Context, params GetEmployeeReq, res *GetEmployeeResp, userId string) error
	AddEmployee(ctx context.Context, req AddEmployeeReq, res *AddEmployeeRes) error
	UpdateEmployee(ctx context.Context, req UpdateEmployeeReq, res *AddEmployeeRes) error
}
