package department

import "github.com/JesseNicholas00/GogoManager/utils/ctxrizz"

type repositoryImpl struct {
	dbRizzer ctxrizz.DbContextRizzer
	statements statements
}

func NewRepository(dbRizzer ctxrizz.DbContextRizzer) Repository {
	return &repositoryImpl{
		dbRizzer: dbRizzer,
		statements: prepareStatements(),
	}
}