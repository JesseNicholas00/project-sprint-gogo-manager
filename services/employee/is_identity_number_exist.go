package employee

import (
	"context"
	"log"

	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
)

func (s *employeeServiceImpl) IsIdentityNumberExist(ctx context.Context, identityNumber string) (bool, error) {
	exists, err := s.repo.IsIdentityNumberExist(ctx, identityNumber)
	if err != nil {
		return false, errorutil.AddCurrentContext(err)
	}

	log.Printf("employeeServiceImpl: Identity number %s exists: %t", identityNumber, exists)
	return exists, nil
}
