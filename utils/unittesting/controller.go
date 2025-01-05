package unittesting

import (
	"github.com/labstack/echo/v4"
)

func CallController(
	ctx echo.Context,
	ctrl func(echo.Context) error,
) {
	err := ctrl(ctx)
	ctx.Echo().HTTPErrorHandler(err, ctx)
}

func CallControllerWithRole(
	ctx echo.Context,
	ctrl func(echo.Context) error,
) {
	err := ctrl(ctx)
	ctx.Echo().HTTPErrorHandler(err, ctx)
}
