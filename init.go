package main

import (
	"github.com/JesseNicholas00/GogoManager/controllers"
	authCtrl "github.com/JesseNicholas00/GogoManager/controllers/auth"
	departmentCtrl "github.com/JesseNicholas00/GogoManager/controllers/department"
	fileCtrl "github.com/JesseNicholas00/GogoManager/controllers/file"
	"github.com/JesseNicholas00/GogoManager/middlewares"
	authRepo "github.com/JesseNicholas00/GogoManager/repos/auth"
	departmentRepo "github.com/JesseNicholas00/GogoManager/repos/department"
	authSvc "github.com/JesseNicholas00/GogoManager/services/auth"
	departmentSvc "github.com/JesseNicholas00/GogoManager/services/department"
	"github.com/JesseNicholas00/GogoManager/utils/ctxrizz"
	"github.com/JesseNicholas00/GogoManager/utils/logging"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/jmoiron/sqlx"
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

	userMw := middlewares.NewAuthMiddleware(authService)

	imageController := fileCtrl.NewImageController(
		uploader,
		cfg.awsS3BucketName,
		userMw,
	)
	ctrls = append(ctrls, imageController)
	departmentRepository := departmentRepo.NewDepartmentRepository(dbRizzer)
	departmentService := departmentSvc.NewDepartmentService(departmentRepository, dbRizzer)
	departmentMw := middlewares.NewAuthMiddleware(authService)
	departmentController := departmentCtrl.NewDepartmentController(departmentService, departmentMw)
	ctrls = append(ctrls, departmentController)

	return
}
