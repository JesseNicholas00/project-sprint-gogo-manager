package department

import (
	"github.com/JesseNicholas00/GogoManager/repos/department"
	"github.com/JesseNicholas00/GogoManager/utils/ctxrizz"
)

type departmentServiceImpl struct {
	repo     department.DepartmentRepository
	dbRizzer ctxrizz.DbContextRizzer
}

func NewDepartmentService(
	repo department.DepartmentRepository,
	dbRizzer ctxrizz.DbContextRizzer,
) DepartmentService {
	return &departmentServiceImpl{
		repo:     repo,
		dbRizzer: dbRizzer,
	}
}
