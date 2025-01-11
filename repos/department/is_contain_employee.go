package department

import (
	"context"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/google/uuid"
)

func (repo *departmentRepositoryImpl) IsContainEmployee(ctx context.Context, departmentID uuid.UUID) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}

	ctx, sess, err := repo.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return false, err
	}

	var count int
	err = sess.Stmt(ctx, repo.statements.isContainEmployee).QueryRowxContext(ctx, departmentID).Scan(&count)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return false, err
	}

	return count > 0, nil
}
