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

	g := server.Group("/v1/employee")
	g.GET("", c.getEmployeeByFilters, c.authMw.Process)
	g.POST("", c.addEmployee, c.authMw.Process)
	g.PATCH("/:identityNumber", c.updateEmployee, c.authMw.Process)

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
