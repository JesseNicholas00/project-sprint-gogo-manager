package department

import (
	"github.com/JesseNicholas00/GogoManager/utils/statementutil"
	"github.com/jmoiron/sqlx"
)

type statements struct {
	add *sqlx.NamedStmt
}

func prepareStatements() statements {
	return statements{
		add: statementutil.MustPrepareNamed(`
			INSERT INTO departments (id, name, manager_id)
			VALUES (:id, :name, :manager_id)
			RETURNING id, name, manager_id;
		`),
	}
}
