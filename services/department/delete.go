package department

import (
	"context"
	"errors"
	"github.com/JesseNicholas00/GogoManager/repos/department"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/JesseNicholas00/GogoManager/utils/transaction"
	"github.com/google/uuid"
)

func (s *serviceImpl) DeleteDepartment(ctx context.Context, req DeleteDepartmentReq, managerId uuid.UUID) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	ctx, sess, err := s.dbRizzer.GetOrAppendTx(ctx)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	return transaction.RunWithAutoCommit(&sess, func() error {
		// TODO: Check for employees

		err = s.repo.DeleteDepartment(ctx, department.Department{
			Id:        req.DepartmentId,
			ManagerId: managerId,
		})
		if err != nil {
			switch {
			case errors.Is(err, department.ErrDepartmentNotFound):
				return ErrDepartmentNotFound
			default:
				return errorutil.AddCurrentContext(err)
			}
		}

		return nil
	})
}
