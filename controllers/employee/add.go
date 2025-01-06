package employee

import (
	"net/http"

	"github.com/JesseNicholas00/GogoManager/services/employee"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/JesseNicholas00/GogoManager/utils/request"
	"github.com/labstack/echo/v4"
)

func (c *employeeController) addEmployee(ctx echo.Context) error {
	req := employee.AddEmployeeReq{}
	if err := request.BindAndValidate(ctx, &req); err != nil {
		return err
	}

	res := employee.AddEmployeeRes{}

	if err := c.service.AddEmployee(ctx.Request().Context(), req, &res); err != nil {
		return errorutil.AddCurrentContext(err)
	}

	return ctx.JSON(http.StatusOK, res)
}
