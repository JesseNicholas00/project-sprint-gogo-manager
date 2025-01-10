package department

import (
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
		return errorutil.AddCurrentContext(err)
	}

	// Get user ID (this is manager ID)
	userId := ctx.Get("session").(auth.GetSessionFromTokenRes).UserId

	managerId, err := uuid.Parse(userId)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	// Validate data
	req := department.AddDepartmentReq{}
	if err := request.BindAndValidate(ctx, &req); err != nil {
		return err
	}

	res := department.AddDepartmentRes{}

	// Get department by ID first
	if err := ctrl.service.GetDepartmentById(ctx.Request().Context(), &res, departmentId, managerId); err != nil {
		return errorutil.AddCurrentContext(err)
	}

	// Update found department
	if err = ctrl.service.UpdateDepartment(ctx.Request().Context(), req, &res, departmentId); err != nil {
		return errorutil.AddCurrentContext(err)
	}

	return ctx.JSON(http.StatusOK, res)
}
