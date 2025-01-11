package department

import (
	"github.com/JesseNicholas00/GogoManager/utils/statementutil"
	"github.com/jmoiron/sqlx"
)

type statements struct {
	add               *sqlx.NamedStmt
	update            *sqlx.NamedStmt
	get               *sqlx.NamedStmt
	getById           *sqlx.NamedStmt
	searchByName      *sqlx.NamedStmt
	delete            *sqlx.NamedStmt
	isContainEmployee *sqlx.Stmt
}

func prepareStatements() statements {
	return statements{
		add: statementutil.MustPrepareNamed(`
			INSERT INTO departments (department_id, name, manager_id)
			VALUES (:department_id, :name, :manager_id);
		`),
		update: statementutil.MustPrepareNamed(`
			UPDATE departments
			SET name = :name
			WHERE department_id = :department_id
		`),
		getById: statementutil.MustPrepareNamed(`
			SELECT department_id, name
			FROM departments
			WHERE 1=1
			AND department_id = :department_id
			AND manager_id = :manager_id;
			`),
		get: statementutil.MustPrepareNamed(`
			SELECT department_id, name 
			FROM departments
			WHERE manager_id = :manager_id
			ORDER BY department_id
			LIMIT :limit OFFSET :offset;
		`),
		searchByName: statementutil.MustPrepareNamed(`
			SELECT department_id, name 
			FROM departments
			WHERE manager_id = :manager_id
			    AND name ILIKE :name
			ORDER BY department_id
			LIMIT :limit OFFSET :offset;
		`),
		delete: statementutil.MustPrepareNamed(`
			DELETE FROM departments 
			WHERE manager_id = :manager_id
			    AND department_id = :department_id;
		`),
		isContainEmployee: statementutil.MustPrepare(`
			SELECT COUNT(1) 
			FROM employees 
			WHERE department_id = $1
		`),
	}
}
