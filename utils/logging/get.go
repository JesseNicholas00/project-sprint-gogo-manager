package logging

import (
	"log/slog"
	"os"
	"strings"

	"github.com/phsym/console-slog"
)

const defaultLogLevel = slog.LevelWarn

var logLevel = defaultLogLevel

var keywordToLevel = map[string]slog.Level{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
}

func SetLogLevel(level string) {
	if lv, ok := keywordToLevel[level]; ok {
		logLevel = lv
	}
}

func GetLogger(context ...string) *slog.Logger {
	logger := slog.New(
		console.NewHandler(
			os.Stderr,
			&console.HandlerOptions{
				Level: logLevel,
			},
		),
	)

	return logger.With("context", strings.Join(context, "."))
}
