package department

import (
	"context"
	"github.com/JesseNicholas00/GogoManager/repos/department"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/google/uuid"
)

func (svc *departmentServiceImpl) GetDepartment(ctx context.Context, params GetDepartmentParams, res *GetDepartmentRes, managerId uuid.UUID) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	departments, err := svc.repo.GetDepartment(ctx, department.FilterDepartment{
		Limit:     *params.Limit,
		Offset:    *params.Offset,
		Name:      params.Name,
		ManagerId: managerId,
	})
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	for _, department := range departments {
		*res = append(*res, AddDepartmentRes{
			DepartmentId: department.Id,
			Name:         department.Name,
		})
	}

	return nil
}
