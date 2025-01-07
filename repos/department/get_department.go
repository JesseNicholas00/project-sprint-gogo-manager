package department

import (
	"context"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
)

func (repo *departmentRepositoryImpl) GetDepartment(ctx context.Context, filter FilterDepartment) ([]Department, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	ctx, sess, err := repo.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return nil, err
	}

	findAll := repo.statements.get
	if filter.Name != "" {
		filter.Name = "%" + filter.Name + "%"
		findAll = repo.statements.searchByName
	}

	rows, err := sess.NamedStmt(ctx, findAll).Queryx(filter)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return nil, err
	}
	defer rows.Close()

	departments := make([]Department, 0, filter.Limit)
	for rows.Next() {
		var department Department
		err = rows.StructScan(&department)
		if err != nil {
			err = errorutil.AddCurrentContext(err)
			return nil, err
		}
		departments = append(departments, department)
	}

	return departments, nil
}
