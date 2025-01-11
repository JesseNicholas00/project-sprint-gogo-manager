package employee

import (
	"github.com/JesseNicholas00/GogoManager/controllers"
	"github.com/JesseNicholas00/GogoManager/middlewares"
	"github.com/JesseNicholas00/GogoManager/services/employee"
	"github.com/labstack/echo/v4"
)

type employeeController struct {
	service employee.EmployeeService
	authMw  middlewares.Middleware
}

func (c *employeeController) Register(server *echo.Echo) error {

	g := server.Group("/v1/employee", c.authMw.Process)
	g.GET("", c.getEmployeeByFilters)
	g.POST("", c.addEmployee)
	g.PATCH("/:identityNumber", c.updateEmployee)
	g.DELETE("/:identityNumber", c.deleteEmployee)

	return nil
}

func NewEmployeeController(
	service employee.EmployeeService,
	authMw middlewares.Middleware,
) controllers.Controller {
	return &employeeController{
		service: service,
		authMw:  authMw,
	}
}
