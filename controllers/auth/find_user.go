package auth

import (
	"net/http"

	"github.com/JesseNicholas00/GogoManager/services/auth"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/labstack/echo/v4"
)

func (ctrl *authController) findUser(c echo.Context) error {
	userId := c.Get("session").(auth.GetSessionFromTokenRes).UserId

	var res auth.FindUserRes

	if err := ctrl.service.FindUser(
		c.Request().Context(),
		userId,
		&res,
	); err != nil {
		return errorutil.AddCurrentContext(err)
	}
	return c.JSON(http.StatusOK, res)
}
