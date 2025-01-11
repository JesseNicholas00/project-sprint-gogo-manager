package employee

import (
	"github.com/JesseNicholas00/GogoManager/utils/statementutil"
	"github.com/jmoiron/sqlx"
)

type statements struct {
	add                   *sqlx.Stmt
	delete                *sqlx.Stmt
	getByIdentityNumber   *sqlx.Stmt
	update                *sqlx.Stmt
	isIdentityNumberExist *sqlx.Stmt
}

func prepareStatements() statements {
	return statements{
		add: statementutil.MustPrepare(`
			INSERT INTO employees (identity_number, name, employee_image_uri, gender, department_id, user_id)
			VALUES ($1, $2, $3, $4, $5, $6)
			RETURNING identity_number, name, employee_image_uri, gender, department_id
		`),
		delete: statementutil.MustPrepare(`
			DELETE FROM employees
			WHERE identity_number = $1 and user_id = $2
		`),
		getByIdentityNumber: statementutil.MustPrepare(`
			SELECT identity_number, name, employee_image_uri, gender, department_id
			FROM employees
			WHERE identity_number = $1 and user_id = $2
		`),
		update: statementutil.MustPrepare(`
			UPDATE employees
			SET identity_number = $1, name = $2, employee_image_uri = $3, gender = $4, department_id = $5
			WHERE identity_number = $6 and user_id = $7
			RETURNING identity_number, name, employee_image_uri, gender, department_id
		`),
		isIdentityNumberExist: statementutil.MustPrepare(`
			SELECT EXISTS (
				SELECT 1
				FROM employees
				WHERE identity_number = $1 and user_id = $2
			)
		`),
	}
}
