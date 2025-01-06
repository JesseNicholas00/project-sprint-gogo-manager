package main

import (
	"github.com/JesseNicholas00/GogoManager/controllers"
	authCtrl "github.com/JesseNicholas00/GogoManager/controllers/auth"
	"github.com/JesseNicholas00/GogoManager/middlewares"
	authRepo "github.com/JesseNicholas00/GogoManager/repos/auth"
	authSvc "github.com/JesseNicholas00/GogoManager/services/auth"
	"github.com/JesseNicholas00/GogoManager/utils/ctxrizz"
	"github.com/JesseNicholas00/GogoManager/utils/logging"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/jmoiron/sqlx"

	employeeCtrl "github.com/JesseNicholas00/GogoManager/controllers/employee"
	employeeRepo "github.com/JesseNicholas00/GogoManager/repos/employee"
	employeeSvc "github.com/JesseNicholas00/GogoManager/services/employee"
)

func initControllers(
	cfg ServerConfig,
	db *sqlx.DB,
	uploader *manager.Uploader,
) (ctrls []controllers.Controller) {
	ctrlInitLogger := logging.GetLogger("main", "init", "controllers")
	defer func() {
		if r := recover(); r != nil {
			// add extra context to help debug potential panic
			ctrlInitLogger.Error("panic while initializing controllers: %s", r)
			panic(r)
		}
	}()

	dbRizzer := ctxrizz.NewDbContextRizzer(db)

	authRepository := authRepo.NewAuthRepository(dbRizzer)
	authService := authSvc.NewAuthService(
		authRepository,
		dbRizzer,
		cfg.jwtSecretKey,
		cfg.bcryptSaltCost,
	)
	authController := authCtrl.NewAuthController(authService)
	ctrls = append(ctrls, authController)

	employeeRepository := employeeRepo.NewRepository(dbRizzer)
	employeeService := employeeSvc.NewEmployeeService(employeeRepository, dbRizzer)
	employeeMw := middlewares.NewAuthMiddleware(authService)
	employeeController := employeeCtrl.NewEmployeeController(employeeService, employeeMw)
	ctrls = append(ctrls, employeeController)

	return
}
