package auth

import (
	"github.com/JesseNicholas00/GogoManager/controllers"
	"github.com/JesseNicholas00/GogoManager/services/auth"
	"github.com/labstack/echo/v4"
)

type authController struct {
	service auth.AuthService
}

func (s *authController) Register(server *echo.Echo) error {
	server.POST("/v1/auth", func(c echo.Context) error {
		return s.authenticateUser(c)
	})
	server.PATCH("/v1/user", func(c echo.Context) error {
		return s.updateUser(c)
	})
	return nil
}

func NewAuthController(
	service auth.AuthService,
) controllers.Controller {
	return &authController{
		service: service,
	}
}
