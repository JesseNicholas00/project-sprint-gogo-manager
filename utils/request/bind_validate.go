package request

import (
	"net/http"

	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/JesseNicholas00/GogoManager/utils/validation"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func BindAndValidate[R any](
	ctx echo.Context,
	req *R,
) error {
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": "invalid request",
		})
	}

	if err := ctx.Validate(req); err != nil {
		if err, ok := err.(validator.ValidationErrors); ok {
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
				"message": validation.ConvertToErrList(err),
			})
		}

		return errorutil.AddCurrentContext(err)
	}

	return nil
}
