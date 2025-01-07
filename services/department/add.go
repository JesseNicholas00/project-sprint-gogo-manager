package department

import (
	"context"
	"github.com/JesseNicholas00/GogoManager/repos/department"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/google/uuid"
)

func (svc *departmentServiceImpl) AddDepartment(ctx context.Context, req AddDepartmentReq, res *AddDepartmentRes, managerId uuid.UUID) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	// No need to add tx here since duplicate department names are allowed
	id, err := uuid.NewV7()
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	err = svc.repo.AddDepartment(ctx, department.Department{
		Id:        id,
		Name:      req.Name,
		ManagerId: managerId,
	})
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	*res = AddDepartmentRes{
		DepartmentId: id,
		Name:         req.Name,
	}

	return nil
}
