package department

import (
	"context"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
)

func (r *departmentRepositoryImpl) AddDepartment(ctx context.Context, department Department) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	ctx, sess, err := r.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return err
	}

	_, err = sess.NamedStmt(ctx, r.statements.add).Exec(department)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return err
	}

	return nil
}
