package department

import (
	"context"
	"github.com/JesseNicholas00/GogoManager/repos/department"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/google/uuid"
)

func (s *serviceImpl) GetDepartment(ctx context.Context, params GetDepartmentParams, res *GetDepartmentRes, managerId uuid.UUID) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	departments, err := s.repo.GetDepartment(ctx, department.FilterDepartment{
		Limit:     *params.Limit,
		Offset:    *params.Offset,
		Name:      params.Name,
		ManagerId: managerId,
	})
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	for _, d := range departments {
		*res = append(*res, AddDepartmentRes{
			DepartmentId: d.Id,
			Name:         d.Name,
		})
	}

	return nil
}
