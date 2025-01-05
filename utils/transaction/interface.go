package transaction

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type DbSession struct {
	Ext       sqlx.ExtContext
	Stmt      func(context.Context, *sqlx.Stmt) *sqlx.Stmt
	NamedStmt func(context.Context, *sqlx.NamedStmt) *sqlx.NamedStmt
	Commit    func() error
	Rollback  func() error
}
