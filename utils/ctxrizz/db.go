package ctxrizz

import (
	"context"
	"database/sql"

	"github.com/JesseNicholas00/GogoManager/utils/transaction"
	"github.com/jmoiron/sqlx"
)

type txContextKeyType struct{}

func (c txContextKeyType) String() string {
	return "txHelper"
}

var txContextKey = txContextKeyType{}

var noop = func() error {
	return nil
}
var noopStmt = func(_ context.Context, s *sqlx.Stmt) *sqlx.Stmt {
	return s
}
var noopNamedStmt = func(_ context.Context, s *sqlx.NamedStmt) *sqlx.NamedStmt {
	return s
}

// ugly workaround for interface{} parameter
func txStmt(tx *sqlx.Tx) func(context.Context, *sqlx.Stmt) *sqlx.Stmt {
	return func(c context.Context, s *sqlx.Stmt) *sqlx.Stmt {
		return tx.StmtxContext(c, s)
	}
}

func txNamedStmt(
	tx *sqlx.Tx,
) func(context.Context, *sqlx.NamedStmt) *sqlx.NamedStmt {
	return tx.NamedStmtContext
}

func (cb *dbContextRizzerImpl) AppendTx(
	ctx context.Context,
) (context.Context, transaction.DbSession, error) {
	tx, err := cb.db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return ctx, transaction.DbSession{}, err
	}

	return context.WithValue(ctx, txContextKey, tx), transaction.DbSession{
		Ext:       tx,
		Stmt:      txStmt(tx),
		NamedStmt: txNamedStmt(tx),
		Commit:    tx.Commit,
		Rollback:  tx.Rollback,
	}, nil
}

func (cb *dbContextRizzerImpl) GetOrAppendTx(
	ctx context.Context,
) (context.Context, transaction.DbSession, error) {
	// already has tx
	if tx, ok := ctx.Value(txContextKey).(*sqlx.Tx); ok {
		// commit and rollback should be done on the creating layer
		// noop here
		return ctx, transaction.DbSession{
			Ext:       tx,
			Stmt:      txStmt(tx),
			NamedStmt: txNamedStmt(tx),
			Commit:    noop,
			Rollback:  noop,
		}, nil
	}

	return cb.AppendTx(ctx)
}

func (cb *dbContextRizzerImpl) GetOrNoTx(
	ctx context.Context,
) (context.Context, transaction.DbSession, error) {
	// already has tx
	if tx, ok := ctx.Value(txContextKey).(*sqlx.Tx); ok {
		// commit and rollback should be done on the creating layer
		// noop here
		return ctx, transaction.DbSession{
			Ext:       tx,
			Stmt:      txStmt(tx),
			NamedStmt: txNamedStmt(tx),
			Commit:    noop,
			Rollback:  noop,
		}, nil
	}

	return ctx, transaction.DbSession{
		Ext:       cb.db,
		Stmt:      noopStmt,
		NamedStmt: noopNamedStmt,
		Commit:    noop,
		Rollback:  noop,
	}, nil
}
