package main

import (
	"context"
	"fmt"

	"github.com/JesseNicholas00/GogoManager/utils/statementutil"

	"github.com/JesseNicholas00/GogoManager/middlewares"
	"github.com/JesseNicholas00/GogoManager/utils/logging"
	"github.com/JesseNicholas00/GogoManager/utils/migration"
	"github.com/JesseNicholas00/GogoManager/utils/validation"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := loadConfig()
	if err != nil {
		logging.GetLogger("config").Error(err.Error())
	}

	logging.SetLogLevel(cfg.logLevel)

	mainInitLogger := logging.GetLogger("main", "init")

	mainInitLogger.Debug(fmt.Sprintf("%+v", cfg))

	if cfg.migrateDownOnStart {
		if err := migration.MigrateDown(cfg.dbString, "migrations"); err != nil {
			mainInitLogger.Error(
				fmt.Sprintf("failed to migrate down db: %s", err),
			)
			return
		}
	}
	if cfg.migrateUpOnStart {
		if err := migration.MigrateUp(cfg.dbString, "migrations"); err != nil {
			mainInitLogger.Error(
				fmt.Sprintf("failed to migrate up db: %s", err),
			)
			return
		}
	}

	db, err := sqlx.Connect("postgres", cfg.dbString)
	if err != nil {
		mainInitLogger.Error(err.Error())
		return
	}

	db.SetMaxOpenConns(cfg.dbMaxOpenConns)
	db.SetMaxIdleConns(cfg.dbMaxIdleConns)
	db.SetConnMaxLifetime(cfg.dbMaxConnLifetime)

	defer func() {
		err := db.Close()
		if err != nil {
			mainInitLogger.Error(err.Error())
		}
	}()

	statementutil.SetUp(db)
	defer statementutil.CleanUp()

	creds := credentials.NewStaticCredentialsProvider(
		cfg.awsAccessKeyID,
		cfg.awsSecretAccessKey,
		"",
	)
	awsCfg, err := awsConfig.LoadDefaultConfig(
		context.TODO(),
		awsConfig.WithCredentialsProvider(creds),
		awsConfig.WithRegion(cfg.awsRegion),
	)
	if err != nil {
		mainInitLogger.Error("panic while initializing aws configs: %s", err)
	}
	client := s3.NewFromConfig(awsCfg)
	uploader := manager.NewUploader(client)

	controllers := initControllers(cfg, db, uploader)

	server := echo.New()

	if cfg.traceSlowEndpoints {
		slowLogger := middlewares.NewSlowTracerMiddleware(cfg.slowThreshold)
		server.Use(slowLogger.Process)
	}

	errorHandler := middlewares.NewLoggingErrorHandlerMiddleware(cfg.showErrorContent)
	server.Use(errorHandler.Process)

	for idx, controller := range controllers {
		if err := controller.Register(server); err != nil {
			msg := fmt.Sprintf(
				"failed during controller registration (%d/%d): %s",
				idx+1,
				len(controllers),
				err,
			)
			mainInitLogger.Error(msg)
			return
		}
	}

	server.Validator = validation.NewEchoValidator()
	server.HideBanner = true

	server.Logger.Fatal(
		server.Start(
			fmt.Sprintf(
				"%s:%d",
				cfg.serverHost,
				cfg.serverPort,
			),
		),
	)
}
