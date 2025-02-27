package department

import (
	"context"
	"github.com/google/uuid"
)

type DepartmentService interface {
	AddDepartment(ctx context.Context, req AddDepartmentReq, res *AddDepartmentRes, managerId uuid.UUID) error
	GetDepartment(ctx context.Context, params GetDepartmentParams, res *GetDepartmentRes, managerId uuid.UUID) error
	GetDepartmentById(ctx context.Context, res *AddDepartmentRes, departmentId uuid.UUID, managerId uuid.UUID) error
	UpdateDepartment(ctx context.Context, req UpdateDepartmentReq, res *AddDepartmentRes, departmentId uuid.UUID, managerId uuid.UUID) error
	DeleteDepartment(ctx context.Context, req DeleteDepartmentReq, managerId uuid.UUID) error
}
