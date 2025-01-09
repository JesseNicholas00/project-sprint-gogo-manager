package auth

import (
	"errors"
	"net/http"

	"github.com/JesseNicholas00/GogoManager/services/auth"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/JesseNicholas00/GogoManager/utils/request"
	"github.com/labstack/echo/v4"
)

func (ctrl *authController) updateUser(c echo.Context) error {
	var req auth.UpdateUserReq
	if err := request.BindAndValidate(c, &req); err != nil {
		return err
	}

	userId := c.Get("session").(auth.GetSessionFromTokenRes).UserId

	var res auth.UpdateUserRes

	if err := ctrl.service.UpdateUser(
		c.Request().Context(),
		userId,
		req,
		&res,
	); err != nil {
		switch {
		case errors.Is(err, auth.ErrEmailAlreadyRegistered):
			return echo.NewHTTPError(http.StatusConflict, echo.Map{
				"message": "email already used by another person",
			})
		default:
			return errorutil.AddCurrentContext(err)
		}
	}
	return c.JSON(http.StatusCreated, res)
}
