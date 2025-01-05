package middlewares

import (
	"log/slog"
	"time"

	"github.com/JesseNicholas00/GogoManager/utils/logging"
	"github.com/labstack/echo/v4"
)

type slowTracerMiddleware struct {
	logger *slog.Logger
	thresh time.Duration
}

func (mw *slowTracerMiddleware) Process(
	next echo.HandlerFunc,
) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		logging.LogIfSlow(mw.thresh, mw.logger, c.Path(), func() {
			err = next(c)
		})
		return
	}
}

func NewSlowTracerMiddleware(slowThreshold time.Duration) Middleware {
	slowTracerMwLogger := logging.GetLogger(
		"slowEndpointTracer",
		slowThreshold.String(),
	)
	return &slowTracerMiddleware{
		logger: slowTracerMwLogger,
		thresh: slowThreshold,
	}
}
