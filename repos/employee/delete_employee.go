package employee

import (
	"context"

	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/google/uuid"
)

func (repo *repositoryEmployeeImpl) DeleteEmployee(ctx context.Context, employee Employee, userId uuid.UUID) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	ctx, sess, err := repo.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	result, err := sess.Stmt(ctx, repo.statements.delete).Exec(employee.IdentityNumber, userId)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	if rowsAffected == 0 {
		err = ErrEmployeeNotFound
		return err
	}

	return nil
}
