package department

import (
	"errors"
	"net/http"

	"github.com/JesseNicholas00/GogoManager/services/auth"
	"github.com/JesseNicholas00/GogoManager/services/department"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/JesseNicholas00/GogoManager/utils/request"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (ctrl *departmentController) updateDepartment(ctx echo.Context) error {
	// Get path variable (departmentId)
	departmentId, err := uuid.Parse(ctx.Param("departmentId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, echo.Map{
			"message": "invalid id",
		})
	}

	// Get user ID (this is manager ID)
	userId := ctx.Get("session").(auth.GetSessionFromTokenRes).UserId

	managerId, err := uuid.Parse(userId)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	// Validate data
	req := department.UpdateDepartmentReq{}
	if err := request.BindAndValidate(ctx, &req); err != nil {
		return err
	}

	res := department.AddDepartmentRes{}

	// Update found department
	if err = ctrl.service.UpdateDepartment(ctx.Request().Context(), req, &res, departmentId, managerId); err != nil {
		switch {
		case errors.Is(err, department.ErrDepartmentNotFound):
			return echo.NewHTTPError(http.StatusNotFound, echo.Map{
				"message": "department not found",
			})
		default:
			return errorutil.AddCurrentContext(err)
		}
	}

	return ctx.JSON(http.StatusOK, res)
}
