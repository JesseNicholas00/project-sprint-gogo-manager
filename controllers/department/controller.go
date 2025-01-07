package department

import (
	"github.com/JesseNicholas00/GogoManager/controllers"
	"github.com/JesseNicholas00/GogoManager/middlewares"
	"github.com/JesseNicholas00/GogoManager/services/department"
	"github.com/labstack/echo/v4"
)

type departmentController struct {
	service department.DepartmentService
	authMw  middlewares.Middleware
}

func (ctrl *departmentController) Register(server *echo.Echo) error {
	g := server.Group("/v1/department")

	g.POST("", ctrl.addDepartment, ctrl.authMw.Process)

	return nil
}

func NewDepartmentController(
	service department.DepartmentService,
	authMw middlewares.Middleware,
) controllers.Controller {
	return &departmentController{
		service: service,
		authMw:  authMw,
	}
}
