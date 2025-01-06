package department

import (
	"errors"
	"github.com/JesseNicholas00/GogoManager/services/auth"
	"github.com/JesseNicholas00/GogoManager/services/department"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/JesseNicholas00/GogoManager/utils/request"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c *controller) deleteDepartment(ctx echo.Context) error {
	req := department.DeleteDepartmentReq{}
	if err := request.BindAndValidate(ctx, &req); err != nil {
		return err
	}

	user, ok := ctx.Get("session").(auth.GetSessionFromTokenRes)
	if !ok {
		return errorutil.AddCurrentContext(nil, "failed to get user from context")
	}

	managerId, err := uuid.Parse(user.UserId)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	if err := c.service.DeleteDepartment(ctx.Request().Context(), req, managerId); err != nil {
		switch {
		case errors.Is(err, department.ErrDepartmentNotFound):
			return echo.NewHTTPError(404, echo.Map{
				"message": "department not found",
			})
		default:
			return errorutil.AddCurrentContext(err)
		}
	}

	return ctx.NoContent(http.StatusOK)
}
