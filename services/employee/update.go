package employee

import (
	"context"
	"errors"
	repoEmployee "github.com/JesseNicholas00/GogoManager/repos/employee"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/JesseNicholas00/GogoManager/utils/transaction"
)

func (svc *employeeServiceImpl) UpdateEmployee(
	ctx context.Context,
	req UpdateEmployeeReq,
	res *AddEmployeeRes,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	ctx, sess, err := svc.dbRizzer.GetOrAppendTx(ctx)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	return transaction.RunWithAutoCommit(&sess, func() error {
		employee, err := svc.repo.FindEmployeeByIdentityNumber(ctx, req.ParamIdentityNumber, req.UserID)

		// Check if employee is not found
		if err != nil {
			switch err {
			case repoEmployee.ErrIdentityNumberNotFound:
				return ErrIdentityNumberNotFound
			default:
				return errorutil.AddCurrentContext(err)

			}
		}

		if req.IdentityNumber.V != nil && *req.IdentityNumber.V != employee.IdentityNumber {
			_, err := svc.repo.FindEmployeeByIdentityNumber(ctx, *req.IdentityNumber.V, req.UserID)
			if err == nil {
				return ErrIdentityNumberAlreadyExists
			}

			if !errors.Is(err, repoEmployee.ErrIdentityNumberNotFound) {
				return errorutil.AddCurrentContext(err)
			}

			employee.IdentityNumber = *req.IdentityNumber.V
		}

		if req.Name.V != nil {
			employee.Name = *req.Name.V
		}

		if req.EmployeeImageUri.V != nil {
			employee.EmployeeImageUri = *req.EmployeeImageUri.V
		}

		if req.Gender.V != nil {
			employee.Gender = *req.Gender.V
		}

		if req.DepartmentId.V != nil {
			employee.DepartmentId = *req.DepartmentId.V
		}

		result, err := svc.repo.UpdateEmployee(ctx, employee, req.ParamIdentityNumber, req.UserID)
		if err != nil {
			return errorutil.AddCurrentContext(err)
		}

		*res = AddEmployeeRes{
			IdentityNumber:   result.IdentityNumber,
			Name:             result.Name,
			EmployeeImageUri: result.EmployeeImageUri,
			Gender:           result.Gender,
			DepartmentId:     result.DepartmentId,
		}

		return nil
	})
}
