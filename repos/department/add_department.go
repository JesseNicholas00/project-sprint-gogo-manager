package department

import (
	"context"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
)

func (r *repositoryImpl) AddDepartment(ctx context.Context, department Department) (Department, error) {
	if err := ctx.Err(); err != nil {
		return Department{}, err
	}

	ctx, sess, err := r.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return Department{}, err
	}

	row := sess.NamedStmt(ctx, r.statements.add).QueryRowx(department)

	var res Department
	err = row.StructScan(&res)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return Department{}, err
	}

	return department, nil
}