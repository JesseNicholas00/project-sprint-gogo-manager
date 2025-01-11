package employee

import (
	"context"

	"github.com/JesseNicholas00/GogoManager/repos/employee"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/google/uuid"
)

func (s *employeeServiceImpl) AddEmployee(ctx context.Context, req AddEmployeeReq, res *AddEmployeeRes, userId uuid.UUID) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	err := s.repo.AddEmployee(ctx, employee.Employee{
		IdentityNumber:   req.IdentityNumber,
		Name:             req.Name,
		EmployeeImageUri: req.EmployeeImageUri,
		Gender:           req.Gender,
		DepartmentId:     req.DepartmentId,
	}, userId)

	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	res.IdentityNumber = req.IdentityNumber
	res.Name = req.Name
	res.EmployeeImageUri = req.EmployeeImageUri
	res.Gender = req.Gender
	res.DepartmentId = req.DepartmentId

	return nil
}
