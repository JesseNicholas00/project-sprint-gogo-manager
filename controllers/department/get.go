package department

import (
	"github.com/JesseNicholas00/GogoManager/services/auth"
	"github.com/JesseNicholas00/GogoManager/services/department"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/JesseNicholas00/GogoManager/utils/request"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (ctrl *departmentController) getDepartment(ctx echo.Context) error {
	req := department.GetDepartmentParams{}
	if err := request.BindAndValidate(ctx, &req); err != nil {
		return err
	}

	if req.Limit == nil || *req.Limit <= 0 {
		req.Limit = new(int)
		*req.Limit = 5
	}

	if req.Offset == nil || *req.Offset < 0 {
		req.Offset = new(int)
		*req.Offset = 0
	}

	userId := ctx.Get("session").(auth.GetSessionFromTokenRes).UserId

	managerId, err := uuid.Parse(userId)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	res := department.GetDepartmentRes{}

	if err := ctrl.service.GetDepartment(ctx.Request().Context(), req, &res, managerId); err != nil {
		return errorutil.AddCurrentContext(err)
	}

	return ctx.JSON(http.StatusOK, res)
}
