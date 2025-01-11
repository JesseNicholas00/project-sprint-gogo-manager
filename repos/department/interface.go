package department

import (
	"context"
	"github.com/google/uuid"
)

type DepartmentRepository interface {
	AddDepartment(ctx context.Context, department Department) error
	UpdateDepartment(ctx context.Context, department Department) error
	GetDepartment(ctx context.Context, filter FilterDepartment) ([]Department, error)
	GetDepartmentById(ctx context.Context, departmentId uuid.UUID, managerId uuid.UUID) (*Department, error)
	DeleteDepartment(ctx context.Context, department Department) error
	IsContainEmployee(ctx context.Context, departmentID uuid.UUID) (bool, error)
}
