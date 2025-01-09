package employee

import (
	"context"

	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
)

func (r *repositoryEmployeeImpl) FindEmployeeByIdentityNumber(
	ctx context.Context,
	identityNumber string,
) (res Employee, err error) {

	if err := ctx.Err(); err != nil {
		return Employee{}, err
	}

	ctx, sess, err := r.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return Employee{}, err
	}

	rows, err := sess.
		Stmt(ctx, r.statements.getByIdentityNumber).
		QueryxContext(ctx, identityNumber)

	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.StructScan(&res)
		if err != nil {
			err = errorutil.AddCurrentContext(err)
			return
		}
	}

	if res.IdentityNumber == "" {
		err = ErrIdentityNumberNotFound
		return
	}

	return
}
