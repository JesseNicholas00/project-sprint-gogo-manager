package employee

import (
	"context"
	"database/sql"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
)

func (r *repositoryEmployeeImpl) IsIdentityNumberExist(ctx context.Context, identityNumber, userID string) (bool, error) {
	var exists bool
	ctx, sess, err := r.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		return false, errorutil.AddCurrentContext(err)
	}

	err = sess.Stmt(ctx, r.statements.isIdentityNumberExist).QueryRowContext(ctx, identityNumber, userID).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		err = errorutil.AddCurrentContext(err)
		return false, err
	}

	return exists, nil

}
