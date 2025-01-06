package department_test

import (
	"github.com/JesseNicholas00/GogoManager/repos/department"
	"github.com/JesseNicholas00/GogoManager/utils/ctxrizz"
	"github.com/JesseNicholas00/GogoManager/utils/unittesting"
	"testing"
)

func NewWithTestDatabase(t *testing.T) department.DepartmentRepository {
	db := unittesting.SetupTestDatabase("../../migrations", t)
	return department.NewDepartmentRepository(ctxrizz.NewDbContextRizzer(db))
}
