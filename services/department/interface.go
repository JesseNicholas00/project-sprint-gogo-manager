package department

import (
	"context"
	"github.com/google/uuid"
)

type DepartmentService interface {
	AddDepartment(ctx context.Context, req AddDepartmentReq, res *AddDepartmentRes, managerId uuid.UUID) error
	GetDepartment(ctx context.Context, params GetDepartmentParams, res *GetDepartmentRes, managerId uuid.UUID) error
	DeleteDepartment(ctx context.Context, req DeleteDepartmentReq, managerId uuid.UUID) error
}
