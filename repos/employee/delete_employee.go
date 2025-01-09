package employee

import (
	"context"

	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
)

func (repo *repositoryEmployeeImpl) DeleteEmployee(ctx context.Context, employee Employee) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	ctx, sess, err := repo.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	result, err := sess.NamedStmt(ctx, repo.statements.delete).Exec(employee)
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
