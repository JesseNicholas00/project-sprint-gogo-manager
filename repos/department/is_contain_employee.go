package department

import (
	"context"
	"github.com/google/uuid"
)

func (repo *departmentRepositoryImpl) IsContainEmployee(ctx context.Context, departmentID uuid.UUID) (bool, error) {
	return true, nil
}
