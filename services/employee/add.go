package employee

import (
	"context"

	"github.com/JesseNicholas00/GogoManager/repos/employee"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
)

func (s *employeeServiceImpl) AddEmployee(ctx context.Context, req AddEmployeeReq, res *AddEmployeeRes) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	created, err := s.repo.AddEmployee(ctx, employee.Employee{
		IdentityNumber:   req.IdentityNumber,
		Name:             req.Name,
		EmployeeImageUri: req.EmployeeImageUri,
		Gender:           req.Gender,
		DepartmentId:     req.DepartmentId,
	})
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	res.IdentityNumber = created.IdentityNumber
	res.EmployeeImageUri = created.EmployeeImageUri
	res.Gender = created.Gender
	res.Name = created.Name
	res.DepartmentId = created.DepartmentId
	return nil
}
