package auth

import (
	"github.com/JesseNicholas00/GogoManager/controllers"
	"github.com/JesseNicholas00/GogoManager/middlewares"
	"github.com/JesseNicholas00/GogoManager/services/auth"
	"github.com/labstack/echo/v4"
)

type authController struct {
	service auth.AuthService
	authMw  middlewares.Middleware
}

func (ctrl *authController) Register(server *echo.Echo) error {
	server.POST("/v1/auth", ctrl.authenticateUser)
	server.PATCH("/v1/user", ctrl.updateUser, ctrl.authMw.Process)
	return nil
}

func NewAuthController(
	service auth.AuthService,
	authMw middlewares.Middleware,
) controllers.Controller {
	return &authController{
		service: service,
		authMw:  authMw,
	}
}
