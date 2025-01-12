package employee

import (
	"context"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
)

func (s *employeeServiceImpl) IsIdentityNumberExist(ctx context.Context, identityNumber, userID string) (bool, error) {
	exists, err := s.repo.IsIdentityNumberExist(ctx, identityNumber, userID)
	if err != nil {
		return false, errorutil.AddCurrentContext(err)
	}

	return exists, nil
}
