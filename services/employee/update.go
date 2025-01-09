package employee

import (
	"context"
	"errors"

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
		employee, err := svc.repo.FindEmployeeByIdentityNumber(ctx, req.ParamIdentityNumber)

		// Check if employee is not found
		if err != nil {
			switch {
			case errors.Is(err, ErrEmployeeNotFound):
				return ErrEmployeeNotFound
			default:
				return errorutil.AddCurrentContext(err)
			}
		}

		if req.IdentityNumber != nil {
			// check if identity number is already exist
			_, err := svc.repo.FindEmployeeByIdentityNumber(ctx, *req.IdentityNumber)
			if err == nil {
				return ErrorEmployeeExist
			}
			employee.IdentityNumber = *req.IdentityNumber
		}

		if req.Name != nil {
			employee.Name = *req.Name
		}

		if req.EmployeeImageUri != nil {
			employee.EmployeeImageUri = *req.EmployeeImageUri
		}

		if req.Gender != nil {
			employee.Gender = *req.Gender
		}

		if req.DepartmentId != nil {
			employee.DepartmentId = *req.DepartmentId
		}

		result, err := svc.repo.UpdateEmployee(ctx, employee, req.ParamIdentityNumber)
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
