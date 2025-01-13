package department

import (
	"context"
	"errors"

	repo "github.com/JesseNicholas00/GogoManager/repos/department"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/google/uuid"
)

func (svc *departmentServiceImpl) UpdateDepartment(ctx context.Context, req UpdateDepartmentReq, res *AddDepartmentRes, departmentId uuid.UUID, managerId uuid.UUID) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	// Get department by ID
	department, err := svc.repo.GetDepartmentById(ctx, departmentId, managerId)
	switch {
	case errors.Is(err, repo.ErrDepartmentNotFound):
		return ErrDepartmentNotFound
	case err != nil:
		return errorutil.AddCurrentContext(err)
	}

	// Update department
	// Return the request if request body is nil
	if req.Name.V == nil {
		res.Name = department.Name
		res.DepartmentId = departmentId

		return nil
	}

	department.Name = *req.Name.V

	err = svc.repo.UpdateDepartment(ctx, *department)

	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	*res = AddDepartmentRes{
		DepartmentId: departmentId,
		Name:         *req.Name.V,
	}

	return nil
}
