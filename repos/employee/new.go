package employee

import "github.com/JesseNicholas00/GogoManager/utils/ctxrizz"

type employeeRepositoryImpl struct {
	dbRizzer ctxrizz.DbContextRizzer
}

func NewRepository(dbRizzer ctxrizz.DbContextRizzer) EmployeeRepository {
	return &employeeRepositoryImpl{
		dbRizzer: dbRizzer,
	}
}
