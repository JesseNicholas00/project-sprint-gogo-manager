package employee_test

import (
	"testing"

	"github.com/JesseNicholas00/GogoManager/repos/employee"
	"github.com/JesseNicholas00/GogoManager/utils/ctxrizz"
	"github.com/JesseNicholas00/GogoManager/utils/unittesting"
)

func NewWithTestDatabase(t *testing.T) employee.EmployeeRepository {
	db := unittesting.SetupTestDatabase("../../migrations", t)
	return employee.NewRepository(ctxrizz.NewDbContextRizzer(db))
}
