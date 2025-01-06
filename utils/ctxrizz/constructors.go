package ctxrizz

import "github.com/jmoiron/sqlx"

type dbContextRizzerImpl struct {
	db *sqlx.DB
}

func NewDbContextRizzer(db *sqlx.DB) DbContextRizzer {
	return &dbContextRizzerImpl{
		db: db,
	}
}
