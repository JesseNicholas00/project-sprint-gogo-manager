package department

import (
	"context"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
)

func (repo *departmentRepositoryImpl) UpdateDepartment(ctx context.Context, department Department) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	ctx, sess, err := repo.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return err
	}

	_, err = sess.NamedStmt(ctx, repo.statements.update).Exec(department)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return err
	}

	return nil
}
