package profile

import (
	"github.com/JesseNicholas00/GogoManager/controllers/auth"
	"github.com/JesseNicholas00/GogoManager/utils/ctxrizz"
)

type profileServiceImpl struct {
	repo     auth.AuthRepository
	dbRizzer ctxrizz.DbContextRizzer
}

func NewProfileService(
	repo auth.AuthRepository,
	dbRizzer ctxrizz.DbContextRizzer,
) ProfileService {
	return &profileServiceImpl{
		repo:     repo,
		dbRizzer: dbRizzer,
	}
}
