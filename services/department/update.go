package department

import (
	"context"

	"github.com/JesseNicholas00/GogoManager/repos/department"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/google/uuid"
)

func (svc *departmentServiceImpl) UpdateDepartment(ctx context.Context, req AddDepartmentReq, res *AddDepartmentRes, departmentId uuid.UUID) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	err := svc.repo.UpdateDepartment(ctx, department.Department{
		Id:   departmentId,
		Name: req.Name,
	})

	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	*res = AddDepartmentRes{
		DepartmentId: departmentId,
		Name:         req.Name,
	}

	return nil
}
