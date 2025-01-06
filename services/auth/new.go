package auth

import (
	"github.com/JesseNicholas00/GogoManager/repos/auth"
	"github.com/JesseNicholas00/GogoManager/utils/ctxrizz"
)

type authServiceImpl struct {
	repo       auth.AuthRepository
	dbRizzer   ctxrizz.DbContextRizzer
	jwtSecret  []byte
	bcryptCost int
}

func NewAuthService(
	repo auth.AuthRepository,
	dbRizzer ctxrizz.DbContextRizzer,
	jwtSecret string,
	bcryptCost int,
) AuthService {
	return &authServiceImpl{
		repo:       repo,
		dbRizzer:   dbRizzer,
		jwtSecret:  []byte(jwtSecret),
		bcryptCost: bcryptCost,
	}
}
