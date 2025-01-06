package employee

import (
	"github.com/JesseNicholas00/GogoManager/controllers"
	"github.com/JesseNicholas00/GogoManager/middlewares"
	"github.com/JesseNicholas00/GogoManager/services/employee"
	"github.com/labstack/echo/v4"
)

type employeeController struct {
	service employee.Service
	authMw  middlewares.Middleware
}

func (c *employeeController) Register(server *echo.Echo) error {
	g := server.Group("/v1/employee")

	g.POST("", c.addEmployee, c.authMw.Process)

	return nil
}

func NewEmployeeController(
	service employee.Service,
	authMw middlewares.Middleware,
) controllers.Controller {
	return &employeeController{
		service: service,
		authMw:  authMw,
	}
}
