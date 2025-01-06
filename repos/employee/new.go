package employee

import "github.com/JesseNicholas00/GogoManager/utils/ctxrizz"

type repositoryEmployeeImpl struct {
	dbRizzer   ctxrizz.DbContextRizzer
	statements statements
}

func NewRepository(dbRizzer ctxrizz.DbContextRizzer) EmployeeRepository {
	return &repositoryEmployeeImpl{
		dbRizzer:   dbRizzer,
		statements: prepareStatements(),
	}
}
