package department

import "context"

type Service interface {
	AddDepartment(ctx context.Context, req AddDepartmentReq, res *AddDepartmentRes) error
}
