package middlewares

import "github.com/labstack/echo/v4"

type Middleware interface {
	Process(next echo.HandlerFunc) echo.HandlerFunc
}
