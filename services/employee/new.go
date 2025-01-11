package employee

import (
	"github.com/JesseNicholas00/GogoManager/repos/employee"
	"github.com/JesseNicholas00/GogoManager/utils/ctxrizz"
)

type employeeServiceImpl struct {
	repo     employee.EmployeeRepository
	dbRizzer ctxrizz.DbContextRizzer
}

func NewEmployeeService(
	repo employee.EmployeeRepository,
	dbRizzer ctxrizz.DbContextRizzer,
) EmployeeService {
	return &employeeServiceImpl{
		repo:     repo,
		dbRizzer: dbRizzer,
	}
}
