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

func (ctrl *departmentController) addDepartment(ctx echo.Context) error {
	req := department.AddDepartmentReq{}
	if err := request.BindAndValidate(ctx, &req); err != nil {
		return err
	}

	userId := ctx.Get("session").(auth.GetSessionFromTokenRes).UserId

	managerId, err := uuid.Parse(userId)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	res := department.AddDepartmentRes{}

	if err := ctrl.service.AddDepartment(ctx.Request().Context(), req, &res, managerId); err != nil {
		return errorutil.AddCurrentContext(err)
	}

	return ctx.JSON(http.StatusCreated, res)
}
