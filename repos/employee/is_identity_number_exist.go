package employee

import (
	"context"
	"database/sql"
	"log"

	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
)

func (r *repositoryEmployeeImpl) IsIdentityNumberExist(ctx context.Context, identityNumber string) (bool, error) {
	var exists bool
	ctx, sess, err := r.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		return false, errorutil.AddCurrentContext(err)
	}

	err = sess.Stmt(ctx, r.statements.isIdentityNumberExist).QueryRowContext(ctx, identityNumber).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("repositoryEmployeeImpl: Identity number %s does not exist", identityNumber)
			return false, nil
		}
		err = errorutil.AddCurrentContext(err)
		return false, err
	}

	log.Printf("repositoryEmployeeImpl: Identity number %s exists: %t", identityNumber, exists)
	return exists, nil

}
