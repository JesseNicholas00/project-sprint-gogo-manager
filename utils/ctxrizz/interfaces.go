package ctxrizz

import (
	"context"

	"github.com/JesseNicholas00/GogoManager/utils/transaction"
)

type DbContextRizzer interface {
	// Starts a new transaction and appends it to the returned context.
	AppendTx(
		ctx context.Context,
	) (context.Context, transaction.DbSession, error)

	// Gets an existing transaction from the given context.
	//
	// Will start a new transaction if there is no existing transaction
	// in the context and append it to the returned context.
	GetOrAppendTx(
		ctx context.Context,
	) (context.Context, transaction.DbSession, error)

	// Gets an existing transaction from the given context.
	//
	// Will NOT start a new transaction if there is no existing transaction
	// in the context.
	GetOrNoTx(
		ctx context.Context,
	) (context.Context, transaction.DbSession, error)
}
