package employee

import (
	"context"

	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/JesseNicholas00/GogoManager/utils/transaction"
	"github.com/google/uuid"
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

		if err != nil {
			return errorutil.AddCurrentContext(err)
		}

		if req.IdentityNumber != "" {
			employee.IdentityNumber = req.IdentityNumber
		}

		if req.Name != "" {
			employee.Name = req.Name
		}

		if req.EmployeeImageUri != "" {
			employee.EmployeeImageUri = req.EmployeeImageUri
		}

		if req.Gender != "" {
			employee.Gender = req.Gender
		}

		if req.DepartmentId != uuid.Nil {
			employee.DepartmentId = req.DepartmentId
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
