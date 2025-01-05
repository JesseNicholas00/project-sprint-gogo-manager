package main

import (
	"fmt"
	"net/url"
	"time"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	serverHost string
	serverPort int

	dbString          string
	dbMaxOpenConns    int
	dbMaxIdleConns    int
	dbMaxConnLifetime time.Duration

	migrateDownOnStart bool
	migrateUpOnStart   bool

	jwtSecretKey   string
	bcryptSaltCost int

	awsAccessKeyID     string
	awsSecretAccessKey string
	awsS3BucketName    string
	awsRegion          string

	logLevel string

	traceSlowEndpoints bool
	slowThreshold      time.Duration

	showErrorContent bool
}

func loadConfig() (cfg ServerConfig, err error) {
	conf := viper.New()
	conf.SetConfigFile(".env")
	conf.SetConfigType("env")
	conf.AutomaticEnv()

	err = conf.ReadInConfig()
	if err != nil {
		return
	}

	conf.SetDefault("HOST", "0.0.0.0")
	conf.SetDefault("PORT", 8080)

	cfg.serverHost = conf.GetString("HOST")
	cfg.serverPort = conf.GetInt("PORT")

	cfg.dbString = fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?%s",
		conf.GetString("DB_USERNAME"),
		// some passwords contain non-safe characters
		url.QueryEscape(conf.GetString("DB_PASSWORD")),
		conf.GetString("DB_HOST"),
		conf.GetInt("DB_PORT"),
		conf.GetString("DB_NAME"),
		conf.GetString("DB_PARAMS"),
	)

	conf.SetDefault("DB_MAX_OPEN_CONNS", 20)
	cfg.dbMaxOpenConns = conf.GetInt("DB_MAX_OPEN_CONNS")

	conf.SetDefault("DB_MAX_IDLE_CONNS", 10)
	cfg.dbMaxIdleConns = conf.GetInt("DB_MAX_IDLE_CONNS")

	conf.SetDefault("DB_MAX_CONN_LIFETIME_MS", 0)
	cfg.dbMaxConnLifetime = time.Duration(
		conf.GetInt64("DB_MAX_CONN_LIFETIME_MS"),
	) * time.Millisecond

	cfg.migrateUpOnStart = conf.GetBool("DB_MIGRATE_UP_ON_START")
	cfg.migrateDownOnStart = conf.GetBool("DB_MIGRATE_DOWN_ON_START")

	cfg.jwtSecretKey = conf.GetString("JWT_SECRET")
	cfg.bcryptSaltCost = conf.GetInt("BCRYPT_SALT")

	cfg.awsAccessKeyID = conf.GetString("AWS_ACCESS_KEY_ID")
	cfg.awsSecretAccessKey = conf.GetString("AWS_SECRET_ACCESS_KEY")
	cfg.awsS3BucketName = conf.GetString("AWS_S3_BUCKET_NAME")
	cfg.awsRegion = conf.GetString("AWS_REGION")

	cfg.logLevel = conf.GetString("LOG_LEVEL")

	cfg.traceSlowEndpoints = conf.GetBool("TRACE_SLOW_ENDPOINTS")

	conf.SetDefault("SLOW_ENDPOINT_THRESHOLD_MS", 100)
	cfg.slowThreshold = time.Duration(
		conf.GetInt64("SLOW_ENDPOINT_THRESHOLD_MS"),
	) * time.Millisecond

	cfg.showErrorContent = conf.GetBool("SHOW_ERROR_CONTENT")

	return
}
