package employee

import (
	"context"

	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
)

func (r *repositoryEmployeeImpl) UpdateEmployee(ctx context.Context, employee Employee, identityNumber string) (res Employee, err error) {
	if err = ctx.Err(); err != nil {
		return
	}

	ctx, sess, err := r.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}

	rows, err := sess.
		Stmt(ctx, r.statements.update).
		QueryxContext(ctx, employee.IdentityNumber,
			employee.Name,
			employee.EmployeeImageUri,
			employee.Gender,
			employee.DepartmentId,
			identityNumber,
		)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}

	err = rows.StructScan(&res)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}

	return employee, nil
}
