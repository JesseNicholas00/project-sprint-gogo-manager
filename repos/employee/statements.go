package employee

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
			INSERT INTO employees (id, identity_number, name, employee_image_uri, gender, department_id)
			VALUES (:id, :identity_number, :name, :employee_image_uri, :gender, :department_id)
			RETURNING id, identity_number, name, employee_image_uri, gender, department_id
		`),
	}
}
