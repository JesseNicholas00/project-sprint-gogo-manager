package department

import (
	"context"
	"github.com/JesseNicholas00/GogoManager/repos/department"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/google/uuid"
)

func (s *serviceImpl) AddDepartment(ctx context.Context, req AddDepartmentReq, res *AddDepartmentRes) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	// No need to add tx here since duplicate department names are allowed
	id, err := uuid.NewV7()
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	created, err := s.repo.AddDepartment(ctx, department.Department{
		Id:   id,
		Name: req.Name,
	})
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	res.DepartmentId = created.Id
	res.Name = created.Name
	return nil
}
