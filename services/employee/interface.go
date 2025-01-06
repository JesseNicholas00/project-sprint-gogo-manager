package employee

import "context"

type Service interface {
	AddEmployee(ctx context.Context, req AddEmployeeReq, res *AddEmployeeRes) error
}
