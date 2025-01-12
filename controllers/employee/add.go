package employee

import (
	"errors"
	"net/http"

	"github.com/JesseNicholas00/GogoManager/services/auth"
	"github.com/JesseNicholas00/GogoManager/services/employee"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/JesseNicholas00/GogoManager/utils/request"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (ctrl *employeeController) addEmployee(ctx echo.Context) error {
	req := employee.AddEmployeeReq{}
	if err := request.BindAndValidate(ctx, &req); err != nil {
		return err
	}

	userId := ctx.Get("session").(auth.GetSessionFromTokenRes).UserId

	managerId, err := uuid.Parse(userId)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	// check identityNumber is exist using IsIdentityNumberExist if exist return error 409
	exists, err := ctrl.service.IsIdentityNumberExist(ctx.Request().Context(), req.IdentityNumber, userId)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}
	if exists {
		return echo.NewHTTPError(409, echo.Map{
			"message": "identity number already exists",
		})
	}

	res := employee.AddEmployeeRes{}

	if err := ctrl.service.AddEmployee(ctx.Request().Context(), req, &res, managerId); err != nil {
		switch {
		case errors.Is(err, employee.ErrDepartmentNotFound):
			return ctx.JSON(http.StatusBadRequest, echo.Map{
				"message": "Department not found",
			})
		default:
			return errorutil.AddCurrentContext(err)
		}
	}

	return ctx.JSON(http.StatusCreated, res)
}
