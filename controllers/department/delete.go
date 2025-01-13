package department

import (
	"errors"
	"github.com/JesseNicholas00/GogoManager/services/auth"
	"github.com/JesseNicholas00/GogoManager/services/department"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (ctrl *departmentController) deleteDepartment(ctx echo.Context) error {
	req := department.DeleteDepartmentReq{}

	var err error
	departmentId := ctx.Param("departmentId")

	req.DepartmentId, err = uuid.Parse(departmentId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, echo.Map{
			"message": "department not found",
		})
	}

	userId := ctx.Get("session").(auth.GetSessionFromTokenRes).UserId

	managerId, err := uuid.Parse(userId)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	if err := ctrl.service.DeleteDepartment(ctx.Request().Context(), req, managerId); err != nil {
		switch {
		case errors.Is(err, department.ErrContainEmployee):
			return echo.NewHTTPError(http.StatusConflict, echo.Map{
				"message": "department has employee(s)",
			})
		case errors.Is(err, department.ErrDepartmentNotFound):
			return echo.NewHTTPError(http.StatusNotFound, echo.Map{
				"message": "department not found",
			})
		default:
			return errorutil.AddCurrentContext(err)
		}
	}

	return ctx.NoContent(http.StatusOK)
}
