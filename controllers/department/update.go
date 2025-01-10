package department

import (
	"github.com/JesseNicholas00/GogoManager/services/auth"
	"github.com/JesseNicholas00/GogoManager/services/department"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (ctrl *departmentController) updateDepartment(ctx echo.Context) error {
	// Get path variable (departmentId)
	departmentId, err := uuid.Parse(ctx.Param("departmentId"))
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	// Get user ID (this is manager ID)
	userId := ctx.Get("session").(auth.GetSessionFromTokenRes).UserId

	managerId, err := uuid.Parse(userId)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	res := department.AddDepartmentRes{}

	if err := ctrl.service.GetDepartmentById(ctx.Request().Context(), &res, departmentId, managerId); err != nil {
		return errorutil.AddCurrentContext(err)
	}

	return ctx.JSON(http.StatusCreated, res)
}
