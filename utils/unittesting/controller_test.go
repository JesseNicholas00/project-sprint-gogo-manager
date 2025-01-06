package unittesting_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JesseNicholas00/GogoManager/utils/unittesting"
	"github.com/labstack/echo/v4"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCallController(t *testing.T) {
	Convey("When wrapping a controller call with uncommited ctx", t, func() {
		ctrl := func(ctx echo.Context) error {
			return echo.NewHTTPError(400, "bruh")
		}

		rec := httptest.NewRecorder()
		ctx := unittesting.CreateEchoContextFromRequest("GET", "/", rec)

		unittesting.CallController(ctx, ctrl)

		Convey("Should commit the ctx", func() {
			So(rec.Code, ShouldEqual, 400)
		})
	})
}
