package employee

import (
	"context"

	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
)

func (r *repositoryEmployeeImpl) AddEmployee(ctx context.Context, employee Employee) (Employee, error) {
	if err := ctx.Err(); err != nil {
		return Employee{}, err
	}

	ctx, sess, err := r.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return Employee{}, err
	}

	row := sess.NamedStmt(ctx, r.statements.add).QueryRowx(employee)

	var res Employee
	err = row.StructScan(&res)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return Employee{}, err
	}

	return employee, nil
}
