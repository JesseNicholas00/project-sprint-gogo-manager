package employee

import (
	"context"
	"errors"

	"github.com/JesseNicholas00/GogoManager/repos/employee"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/JesseNicholas00/GogoManager/utils/transaction"
	"github.com/google/uuid"
)

func (s *employeeServiceImpl) DeleteEmployee(ctx context.Context, req DeleteEmployeeReq, userId uuid.UUID) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	ctx, sess, err := s.dbRizzer.GetOrAppendTx(ctx)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	return transaction.RunWithAutoCommit(&sess, func() error {
		err = s.repo.DeleteEmployee(ctx, employee.Employee{
			IdentityNumber: req.IdentityNumber,
		}, userId)
		if err != nil {
			switch {
			case errors.Is(err, employee.ErrEmployeeNotFound):
				return ErrEmployeeNotFound
			default:
				return errorutil.AddCurrentContext(err)
			}
		}

		return nil
	})
}
