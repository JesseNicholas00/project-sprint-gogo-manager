package logging

import (
	"log/slog"
	"time"
)

func LogIfSlow(
	threshold time.Duration,
	logger *slog.Logger,
	message string,
	op func(),
) {
	startTime := time.Now()
	op()
	timeTaken := time.Since(startTime)

	if timeTaken >= threshold {
		logger.Warn(
			"slow log",
			"message",
			message,
			"duration",
			timeTaken.String(),
		)
	}
}
