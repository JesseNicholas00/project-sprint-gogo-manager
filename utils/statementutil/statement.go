package statementutil

import "github.com/jmoiron/sqlx"

var globalDb *sqlx.DB
var cleanupFuncs []func()

func MustPrepareNamed(sqlQuery string) *sqlx.NamedStmt {
	checkGlobalDb()

	stmt, err := globalDb.PrepareNamed(sqlQuery)
	if err != nil {
		panic(err)
	}
	cleanupFuncs = append(cleanupFuncs, func() {
		stmt.Close()
	})
	return stmt
}

func MustPrepare(sqlQuery string) *sqlx.Stmt {
	checkGlobalDb()

	stmt, err := globalDb.Preparex(sqlQuery)
	if err != nil {
		panic(err)
	}
	cleanupFuncs = append(cleanupFuncs, func() {
		stmt.Close()
	})
	return stmt
}

func SetUp(db *sqlx.DB) {
	globalDb = db
}

func CleanUp() {
	for _, cleanupFunc := range cleanupFuncs {
		cleanupFunc()
	}
}

func checkGlobalDb() {
	if globalDb == nil {
		panic(
			"global db is not set! did you forget to call statementutil.SetUp(db)?",
		)
	}
}
