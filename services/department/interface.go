package department

import (
	"context"
	"github.com/google/uuid"
)

type DepartmentService interface {
	AddDepartment(ctx context.Context, req AddDepartmentReq, res *AddDepartmentRes, managerId uuid.UUID) error
}
