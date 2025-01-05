package auth

import (
	"github.com/JesseNicholas00/GogoManager/utils/ctxrizz"
)

type authRepositoryImpl struct {
	dbRizzer   ctxrizz.DbContextRizzer
	statements statements
}

func NewAuthRepository(dbRizzer ctxrizz.DbContextRizzer) AuthRepository {
	return &authRepositoryImpl{
		dbRizzer:   dbRizzer,
		statements: prepareStatements(),
	}
}
