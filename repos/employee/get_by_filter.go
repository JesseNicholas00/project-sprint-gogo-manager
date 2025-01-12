package employee

import (
	"context"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/JesseNicholas00/GogoManager/utils/mewsql"
)

func (r *repositoryEmployeeImpl) GetEmployeeByFilters(ctx context.Context, filter FilterEmployee) ([]Employee, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	var conditions []mewsql.Condition

	if filter.IdentityNumber != "" {
		conditions = append(conditions,
			mewsql.WithCondition("identity_number ILIKE ?", filter.IdentityNumber+"%"),
		)
	}

	if filter.Name != "" {
		conditions = append(conditions,
			mewsql.WithCondition("name ILIKE ?", "%"+filter.Name+"%"),
		)
	}

	if filter.Gender != "" {
		conditions = append(conditions,
			mewsql.WithCondition("gender = ?", filter.Gender),
		)
	}

	if filter.DepartementId != "" {
		conditions = append(conditions,
			mewsql.WithCondition("department_id = ?", filter.DepartementId),
		)
	}

	conditions = append(conditions, mewsql.WithCondition("user_id = ?", filter.UserId))

	options := []mewsql.SelectOption{
		mewsql.WithLimit(filter.Limit),
		mewsql.WithOffset(filter.Offset),
		mewsql.WithWhere(conditions...),
	}

	sql, args := mewsql.Select(
		`identity_number, name, employee_image_uri, gender, department_id`,
		"employees",
		options...,
	)

	ctx, sess, err := r.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return nil, err
	}

	rows, err := sess.Ext.QueryxContext(ctx, sql, args...)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return nil, err
	}

	defer rows.Close()

	employees := make([]Employee, 0, filter.Limit)
	for rows.Next() {
		var employee Employee
		err = rows.StructScan(&employee)
		if err != nil {
			err = errorutil.AddCurrentContext(err)
			return nil, err
		}
		employees = append(employees, employee)
	}

	return employees, nil
}
