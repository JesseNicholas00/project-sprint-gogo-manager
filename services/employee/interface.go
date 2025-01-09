package employee

import "context"

type Service interface {
	AddEmployee(ctx context.Context, req AddEmployeeReq, res *AddEmployeeRes) error
	DeleteEmployee(ctx context.Context, req DeleteEmployeeReq) error
}
