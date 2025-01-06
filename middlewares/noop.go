package middlewares

import "github.com/labstack/echo/v4"

type noopMiddleware struct {
}

func (*noopMiddleware) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return next
}

func NewNoopMiddleware() Middleware {
	return &noopMiddleware{}
}
