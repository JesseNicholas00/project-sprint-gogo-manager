package employee

import "context"

type EmployeeService interface {
	GetEmployeeByFilters(ctx context.Context, params GetEmployeeReq, res *[]GetEmployeeResp, userId string) error
}
