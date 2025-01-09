package employee

import (
	"errors"
	"net/http"

	"github.com/JesseNicholas00/GogoManager/services/employee"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/JesseNicholas00/GogoManager/utils/request"
	"github.com/labstack/echo/v4"
)

func (c *employeeController) deleteEmployee(ctx echo.Context) error {
	req := employee.DeleteEmployeeReq{}
	if err := request.BindAndValidate(ctx, &req); err != nil {
		return err
	}

	if err := c.service.DeleteEmployee(ctx.Request().Context(), req); err != nil {
		switch {
		case errors.Is(err, employee.ErrEmployeeNotFound):
			return echo.NewHTTPError(404, echo.Map{
				"message": "employee not found",
			})
		default:
			return errorutil.AddCurrentContext(err)
		}
	}

	return ctx.NoContent(http.StatusOK)
}
