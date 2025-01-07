package employee

import (
	"github.com/JesseNicholas00/GogoManager/utils/statementutil"
	"github.com/jmoiron/sqlx"
)

type statements struct {
	add                 *sqlx.NamedStmt
	getByIdentityNumber *sqlx.Stmt
	update              *sqlx.Stmt
}

func prepareStatements() statements {
	return statements{
		add: statementutil.MustPrepareNamed(`
			INSERT INTO employees (identity_number, name, employee_image_uri, gender, department_id)
			VALUES (:identity_number, :name, :employee_image_uri, :gender, :department_id)
			RETURNING identity_number, name, employee_image_uri, gender, department_id
		`),
		getByIdentityNumber: statementutil.MustPrepare(`
			SELECT *
			FROM employees
			WHERE identity_number = $1
		`),
		update: statementutil.MustPrepare(`
			UPDATE employees
			SET identity_number = $1, name = $2, employee_image_uri = $3, gender = $4, department_id = $5
			WHERE identity_number = $6
			RETURNING identity_number, name, employee_image_uri, gender, department_id
		`),
	}
}
