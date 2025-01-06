package department

import (
	"github.com/JesseNicholas00/GogoManager/repos/department"
	"github.com/JesseNicholas00/GogoManager/utils/ctxrizz"
)

type serviceImpl struct {
	repo department.Repository
	dbRizzer ctxrizz.DbContextRizzer
}

func NewService(
	repo department.Repository,
	dbRizzer ctxrizz.DbContextRizzer,
) Service {
	return &serviceImpl{
		repo: repo,
		dbRizzer: dbRizzer,
	}
}