package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/JesseNicholas00/GogoManager/services/auth"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/labstack/echo/v4"
)

type authMiddleware struct {
	service auth.AuthService
	binder  *echo.DefaultBinder
}

func (mw *authMiddleware) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := struct {
			Bearer string `header:"Authorization"`
		}{}

		if err := mw.binder.BindHeaders(c, &header); err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, echo.Map{
				"message": "invalid request",
			})
		}

		if header.Bearer == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, echo.Map{
				"message": "missing bearer token",
			})
		}

		splitByBearer := strings.Split(header.Bearer, "Bearer ")
		if len(splitByBearer) != 2 {
			return echo.NewHTTPError(http.StatusUnauthorized, echo.Map{
				"message": "malformed bearer token",
			})
		}

		token := splitByBearer[1]
		req := auth.GetSessionFromTokenReq{
			AccessToken: token,
		}
		var res auth.GetSessionFromTokenRes
		if err := mw.service.GetSessionFromToken(
			c.Request().Context(),
			req,
			&res,
		); err != nil {
			switch {
			case errors.Is(err, auth.ErrTokenInvalid):
				return echo.NewHTTPError(http.StatusUnauthorized, echo.Map{
					"message": "malformed bearer token",
				})

			case errors.Is(err, auth.ErrTokenExpired):
				return echo.NewHTTPError(http.StatusUnauthorized, echo.Map{
					"message": "expired bearer token",
				})

			default:
				return errorutil.AddCurrentContext(err)
			}
		}

		c.Set("session", res)

		return next(c)
	}
}

func NewAuthMiddleware(
	service auth.AuthService,
) Middleware {
	return &authMiddleware{
		service: service,
		binder:  &echo.DefaultBinder{},
	}
}
