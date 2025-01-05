package auth_test

import (
	"testing"

	"github.com/JesseNicholas00/GogoManager/repos/auth"
	"github.com/JesseNicholas00/GogoManager/utils/ctxrizz"
	"github.com/JesseNicholas00/GogoManager/utils/unittesting"
)

func NewWithTestDatabase(t *testing.T) auth.AuthRepository {
	db := unittesting.SetupTestDatabase("../../migrations", t)
	return auth.NewAuthRepository(ctxrizz.NewDbContextRizzer(db))
}
