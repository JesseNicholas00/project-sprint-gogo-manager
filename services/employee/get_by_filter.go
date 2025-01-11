package employee

import (
	"context"

	"github.com/JesseNicholas00/GogoManager/repos/employee"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
)

func (e *employeeServiceImpl) GetEmployeeByFilters(ctx context.Context, params GetEmployeeReq, res *GetEmployeeResp, userId string) error {

	if err := ctx.Err(); err != nil {
		return err
	}

	employees, err := e.repo.GetEmployeeByFilters(ctx, employee.FilterEmployee{
		Limit:         *params.Limit,
		Offset:        *params.Offset,
		Name:          params.Name,
		Gender:        params.Gender,
		DepartementId: params.DepartementId,
		UserId:        userId,
	})

	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	for _, e := range employees {
		*res = append(*res, AddEmployeeRes{
			IdentityNumber:   e.IdentityNumber,
			Name:             e.Name,
			EmployeeImageUri: e.EmployeeImageUri,
			Gender:           e.Gender,
			DepartmentId:     e.DepartmentId,
		})
	}

	return nil
}
