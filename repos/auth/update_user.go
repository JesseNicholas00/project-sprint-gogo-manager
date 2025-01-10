package auth

import (
	"context"

	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
)

func (repo *authRepositoryImpl) UpdateUser(
	ctx context.Context,
	user User,
) (res User, err error) {
	if err = ctx.Err(); err != nil {
		return
	}

	ctx, sess, err := repo.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}

	rows, err := sess.NamedStmt(ctx, repo.statements.update).Queryx(user)
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

	return
}
