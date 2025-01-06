package department

import (
	"github.com/JesseNicholas00/GogoManager/services/department"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/JesseNicholas00/GogoManager/utils/request"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c *controller) getDepartment(ctx echo.Context) error {
	req := department.GetDepartmentParams{}
	if err := request.BindAndValidate(ctx, &req); err != nil {
		return err
	}

	if req.Limit == nil {
		req.Limit = new(uint)
		*req.Limit = 5
	}

	if req.Offset == nil {
		req.Offset = new(uint)
		*req.Offset = 0
	}

	res := department.GetDepartmentRes{}

	if err := c.service.GetDepartment(ctx.Request().Context(), req, &res); err != nil {
		return errorutil.AddCurrentContext(err)
	}

	return ctx.JSON(http.StatusOK, res)
}
