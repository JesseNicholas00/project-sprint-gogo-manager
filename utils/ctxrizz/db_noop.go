package ctxrizz

import (
	"context"

	"github.com/JesseNicholas00/GogoManager/utils/transaction"
)

type dbContextRizzerNoopImpl struct {
}

// AppendTx implements DbContextRizzer.
func (d *dbContextRizzerNoopImpl) AppendTx(
	ctx context.Context,
) (context.Context, transaction.DbSession, error) {
	return ctx, getDummySession(), nil
}

// GetOrAppendTx implements DbContextRizzer.
func (d *dbContextRizzerNoopImpl) GetOrAppendTx(
	ctx context.Context,
) (context.Context, transaction.DbSession, error) {
	return ctx, getDummySession(), nil
}

// GetOrNoTx implements DbContextRizzer.
func (d *dbContextRizzerNoopImpl) GetOrNoTx(
	ctx context.Context,
) (context.Context, transaction.DbSession, error) {
	return ctx, getDummySession(), nil
}

func NewDbContextNoopRizzer() DbContextRizzer {
	return &dbContextRizzerNoopImpl{}
}

func getDummySession() transaction.DbSession {
	return transaction.DbSession{
		Stmt:      noopStmt,
		NamedStmt: noopNamedStmt,
		Commit:    noop,
		Rollback:  noop,
	}
}
