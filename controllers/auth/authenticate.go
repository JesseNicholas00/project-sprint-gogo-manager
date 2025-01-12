package auth

import (
	"errors"
	"net/http"

	"github.com/JesseNicholas00/GogoManager/services/auth"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/JesseNicholas00/GogoManager/utils/request"
	"github.com/labstack/echo/v4"
)

func (ctrl *authController) authenticateUser(c echo.Context) error {
	var req auth.AuthenticateUserReq
	if err := request.BindAndValidate(c, &req); err != nil {
		return err
	}

	var res auth.AuthenticateUserRes

	if req.Action == "create" {
		if err := ctrl.service.RegisterUser(
			c.Request().Context(),
			req,
			&res,
		); err != nil {
			switch {
			case errors.Is(err, auth.ErrEmailAlreadyRegistered):
				return echo.NewHTTPError(http.StatusConflict, echo.Map{
					"message": "user already exists",
				})
			default:
				return errorutil.AddCurrentContext(err)
			}
		}
		return c.JSON(http.StatusCreated, res)
	}
	if err := ctrl.service.LoginUser(
		c.Request().Context(),
		req,
		&res,
	); err != nil {
		switch {
		case errors.Is(err, auth.ErrUserNotFound):
			return echo.NewHTTPError(http.StatusNotFound, echo.Map{
				"message": "user not found",
			})
		default:
			return errorutil.AddCurrentContext(err)
		}
	}
	return c.JSON(http.StatusOK, res)
}
