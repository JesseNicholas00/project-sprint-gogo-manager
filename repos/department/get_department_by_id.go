package department

import (
	"context"

	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/google/uuid"
)

func (repo *departmentRepositoryImpl) GetDepartmentById(ctx context.Context, departmentId uuid.UUID, managerId uuid.UUID) (*Department, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	ctx, sess, err := repo.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return nil, err
	}

	findById := repo.statements.getById

	// Create map for querying
	filter := map[string]interface{}{
		"department_id": departmentId.String(),
		"manager_id":    managerId.String(),
	}

	var department Department

	err = sess.NamedStmt(ctx, findById).QueryRowx(filter).StructScan(&department)

	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return nil, err
	}

	return &department, nil
}
