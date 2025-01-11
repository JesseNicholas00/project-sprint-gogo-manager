package auth

import (
	"github.com/JesseNicholas00/GogoManager/utils/statementutil"
	"github.com/jmoiron/sqlx"
)

type statements struct {
	create       *sqlx.NamedStmt
	update       *sqlx.NamedStmt
	findByUserId *sqlx.Stmt
	findByEmail  *sqlx.Stmt
}

func prepareStatements() statements {
	return statements{
		create: statementutil.MustPrepareNamed(`
			INSERT INTO users(
				user_id,
				email,
				password
			) VALUES (
				:user_id,
				:email,
				:password
			) RETURNING
				user_id,
				email,
				password
		`),
		update: statementutil.MustPrepareNamed(`
			UPDATE users
			SET
				email = :email,
				user_name = :user_name,
				user_image_uri = :user_image_uri,
				company_name = :company_name,
				company_image_uri = :company_image_uri
			WHERE user_id = :user_id
			RETURNING
				user_id,
				user_name,
				email,
				password,
				user_image_uri,
				company_name,
				company_image_uri
		`),
		findByUserId: statementutil.MustPrepare(`
			SELECT
				*
			FROM
				users
			WHERE
				user_id = $1
		`),
		findByEmail: statementutil.MustPrepare(`
			SELECT
				*
			FROM
				users
			WHERE
				email = $1
		`),
	}
}
