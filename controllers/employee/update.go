package employee

import (
	"errors"
	"net/http"

	"github.com/JesseNicholas00/GogoManager/services/auth"
	"github.com/JesseNicholas00/GogoManager/services/employee"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/JesseNicholas00/GogoManager/utils/request"
	"github.com/labstack/echo/v4"
)

func (c *employeeController) updateEmployee(ctx echo.Context) error {
	req := employee.UpdateEmployeeReq{}
	if err := request.BindAndValidate(ctx, &req); err != nil {
		return err
	}

	req.ParamIdentityNumber = ctx.Param("identityNumber")
	req.UserID = ctx.Get("session").(auth.GetSessionFromTokenRes).UserId

	res := employee.AddEmployeeRes{}

	if err := c.service.UpdateEmployee(ctx.Request().Context(), req, &res); err != nil {
		switch {
		case errors.Is(err, employee.ErrIdentityNumberNotFound):
			return echo.NewHTTPError(http.StatusNotFound, echo.Map{
				"message": "identity number not found",
			})
		case errors.Is(err, employee.ErrIdentityNumberAlreadyExists):
			return echo.NewHTTPError(http.StatusConflict, echo.Map{
				"message": "identity number already exists",
			})
		default:
			return errorutil.AddCurrentContext(err)
		}
	}

	return ctx.JSON(http.StatusOK, res)
}
