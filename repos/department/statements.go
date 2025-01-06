package department

import (
	"github.com/JesseNicholas00/GogoManager/utils/statementutil"
	"github.com/jmoiron/sqlx"
)

type statements struct {
	add          *sqlx.NamedStmt
	get          *sqlx.NamedStmt
	searchByName *sqlx.NamedStmt
}

func prepareStatements() statements {
	return statements{
		add: statementutil.MustPrepareNamed(`
			INSERT INTO departments (department_id, name, manager_id)
			VALUES (:department_id, :name, :manager_id);
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
		get: statementutil.MustPrepareNamed(`
			SELECT id, name 
			FROM departments
			WHERE manager_id = :manager_id
			ORDER BY id
			LIMIT :limit OFFSET :offset;
		`),
		searchByName: statementutil.MustPrepareNamed(`
			SELECT id, name 
			FROM departments
			WHERE manager_id = :manager_id
			    AND name ILIKE :name
			ORDER BY id
			LIMIT :limit OFFSET :offset;
		`),
	}
}
