package transaction

import (
	"database/sql"

	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/pkg/errors"
)

func RunWithAutoCommit(
	sess *DbSession,
	operation func() error,
) error {
	defer func() {
		// must rollback on panic!
		if r := recover(); r != nil {
			err := sess.Rollback()
			if err != nil {
				err = errors.Wrapf(err, "%v", r)
				panic(errorutil.AddCurrentContext(err))
			}
			panic(r)
		}
	}()

	if err := operation(); err != nil {
		errRollback := sess.Rollback()
		if errRollback != nil && !errors.Is(errRollback, sql.ErrTxDone) {
			return errorutil.AddCurrentContext(err)
		}

		return err
	}

	errCommit := sess.Commit()
	if errCommit != nil && !errors.Is(errCommit, sql.ErrTxDone) {
		return errorutil.AddCurrentContext(errCommit)
	}

	return nil
}
