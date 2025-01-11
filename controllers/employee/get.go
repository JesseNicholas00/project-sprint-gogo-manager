package employee

import (
	"net/http"

	"github.com/JesseNicholas00/GogoManager/services/auth"
	"github.com/JesseNicholas00/GogoManager/services/employee"
	"github.com/JesseNicholas00/GogoManager/utils/request"
	"github.com/labstack/echo/v4"
)

func (e *employeeController) getEmployeeByFilters(ctx echo.Context) error {
	var params employee.GetEmployeeReq
	if err := request.BindAndValidate(ctx, &params); err != nil {
		return err
	}

	if params.Limit == nil || *params.Limit < 1 {
		params.Limit = new(int)
		*params.Limit = 5
	}

	if params.Offset == nil || *params.Offset < 1 {
		params.Offset = new(int)
		*params.Offset = 0
	}

	user := ctx.Get("session").(auth.GetSessionFromTokenRes)

	var res employee.GetEmployeeResp
	if err := e.service.GetEmployeeByFilters(ctx.Request().Context(), params, &res, user.UserId); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}
