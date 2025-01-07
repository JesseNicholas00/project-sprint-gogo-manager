package profile

import (
	"github.com/JesseNicholas00/GogoManager/controllers"
	"github.com/JesseNicholas00/GogoManager/services/profile"
	"github.com/labstack/echo/v4"
)

type profileController struct {
	service profile.ProfileService
}

func (s *profileController) Register(server *echo.Echo) error {
	server.PATCH("/v1/auth", func(c echo.Context) error {
		return s.upsertUser(c)
	})
	return nil
}

func NewProfileController(
	service profile.ProfileService,
) controllers.Controller {
	return &profileController{
		service: service,
	}
}
