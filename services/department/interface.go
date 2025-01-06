package department

import (
	"context"
	"github.com/google/uuid"
)

type Service interface {
	AddDepartment(ctx context.Context, req AddDepartmentReq, res *AddDepartmentRes, managerId uuid.UUID) error
}
