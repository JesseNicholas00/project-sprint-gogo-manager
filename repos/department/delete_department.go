package department

import (
	"context"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
)

func (repo *departmentRepositoryImpl) DeleteDepartment(ctx context.Context, department Department) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	ctx, sess, err := repo.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	result, err := sess.NamedStmt(ctx, repo.statements.delete).Exec(department)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	if rowsAffected == 0 {
		err = ErrDepartmentNotFound
		return err
	}

	return nil
}
