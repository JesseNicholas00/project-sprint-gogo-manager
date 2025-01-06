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

func (c *controller) addDepartment(ctx echo.Context) error {
	req := department.AddDepartmentReq{}
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

	res := department.AddDepartmentRes{}

	if err := c.service.AddDepartment(ctx.Request().Context(), req, &res, managerId); err != nil {
		return errorutil.AddCurrentContext(err)
	}

	return ctx.JSON(http.StatusOK, res)
}
