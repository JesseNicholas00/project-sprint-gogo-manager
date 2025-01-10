package department

import (
	"context"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/google/uuid"
)

func (svc *departmentServiceImpl) GetDepartmentById(ctx context.Context, res *AddDepartmentRes, departmentId uuid.UUID, managerId uuid.UUID) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	department, err := svc.repo.GetDepartmentById(ctx, departmentId, managerId)

	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	// Map department to response
	res.Name = department.Name
	res.DepartmentId = departmentId

	return nil
}
