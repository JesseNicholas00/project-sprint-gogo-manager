package employee

import (
	"github.com/JesseNicholas00/GogoManager/utils/statementutil"
	"github.com/jmoiron/sqlx"
)

type statements struct {
	add    *sqlx.NamedStmt
	delete *sqlx.NamedStmt
}

func prepareStatements() statements {
	return statements{
		add: statementutil.MustPrepareNamed(`
			INSERT INTO employees (identity_number, name, employee_image_uri, gender, department_id)
			VALUES (:identity_number, :name, :employee_image_uri, :gender, :department_id)
			RETURNING identity_number, name, employee_image_uri, gender, department_id
		`),
		delete: statementutil.MustPrepareNamed(`
			DELETE FROM employees
			WHERE identity_number = :identity_number
		`),
	}
}
