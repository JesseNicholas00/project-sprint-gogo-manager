package department_test

import (
	"github.com/JesseNicholas00/GogoManager/repos/department"
	"github.com/JesseNicholas00/GogoManager/utils/ctxrizz"
	"github.com/JesseNicholas00/GogoManager/utils/unittesting"
	"testing"
)

func NewWithTestDatabase(t *testing.T) department.Repository {
	db := unittesting.SetupTestDatabase("../../migrations", t)
	return department.NewRepository(ctxrizz.NewDbContextRizzer(db))
}
