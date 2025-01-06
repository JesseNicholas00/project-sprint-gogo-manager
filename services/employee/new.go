package employee

import (
	"github.com/JesseNicholas00/GogoManager/repos/employee"
	"github.com/JesseNicholas00/GogoManager/utils/ctxrizz"
)

type serviceImpl struct {
	repo     employee.EmployeeRepository
	dbRizzer ctxrizz.DbContextRizzer
}

func NewService(
	repo employee.EmployeeRepository,
	dbRizzer ctxrizz.DbContextRizzer,
) Service {
	return &serviceImpl{
		repo:     repo,
		dbRizzer: dbRizzer,
	}
}
