package employee

import (
	"net/http"

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

	res := employee.AddEmployeeRes{}

	if err := c.service.UpdateEmployee(ctx.Request().Context(), req, &res); err != nil {
		return errorutil.AddCurrentContext(err)
	}

	return ctx.JSON(http.StatusOK, res)
}
