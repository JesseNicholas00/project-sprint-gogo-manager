package department

import "github.com/JesseNicholas00/GogoManager/utils/ctxrizz"

type departmentRepositoryImpl struct {
	dbRizzer   ctxrizz.DbContextRizzer
	statements statements
}

func NewDepartmentRepository(dbRizzer ctxrizz.DbContextRizzer) DepartmentRepository {
	return &departmentRepositoryImpl{
		dbRizzer:   dbRizzer,
		statements: prepareStatements(),
	}
}
