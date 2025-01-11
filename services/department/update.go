package department

import (
	"context"

	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/google/uuid"
)

func (svc *departmentServiceImpl) UpdateDepartment(ctx context.Context, req AddDepartmentReq, res *AddDepartmentRes, departmentId uuid.UUID, managerId uuid.UUID) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	// Get department by ID
	department, err := svc.repo.GetDepartmentById(ctx, departmentId, managerId)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	// Update department
	department.Name = req.Name

	err = svc.repo.UpdateDepartment(ctx, *department)

	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	*res = AddDepartmentRes{
		DepartmentId: departmentId,
		Name:         req.Name,
	}

	return nil
}
