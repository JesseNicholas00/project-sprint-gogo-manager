package employee

import (
	"context"

	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/google/uuid"
)

func (r *repositoryEmployeeImpl) AddEmployee(ctx context.Context, employee Employee, userId uuid.UUID) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	ctx, sess, err := r.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return err
	}

	row := sess.NamedStmt(ctx, r.statements.add).QueryRowx(employee)

	var res Employee
	err = row.StructScan(&res)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return err
	}

	return nil
}
