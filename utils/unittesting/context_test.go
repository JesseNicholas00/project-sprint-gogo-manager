package unittesting_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/JesseNicholas00/GogoManager/utils/unittesting"
)

func TestCreateEchoContextFromRequest(t *testing.T) {
	Convey(
		"When creating ctx from a request to `/test`",
		t,
		func() {
			testUrl := "/test"
			Convey("And the request is a GET", func() {
				ctx := unittesting.CreateEchoContextFromRequest(
					http.MethodGet,
					testUrl,
					&httptest.ResponseRecorder{},
				)
				Convey("The created context should be a GET", func() {
					So(ctx.Request().Method, ShouldEqual, http.MethodGet)
				})
				Convey(
					"The created context should have the path `/test/`",
					func() {
						So(ctx.Path(), ShouldEqual, testUrl)
					},
				)
			})

			Convey("And the request is a POST", func() {
				ctx := unittesting.CreateEchoContextFromRequest(
					http.MethodPost,
					testUrl,
					&httptest.ResponseRecorder{},
				)
				Convey("The created context should be a POST", func() {
					So(ctx.Request().Method, ShouldEqual, http.MethodPost)
				})
				Convey(
					"The created context should have the path `/test/`",
					func() {
						So(ctx.Path(), ShouldEqual, testUrl)
					},
				)
			})
		},
	)
}

func TestWithPathParams(t *testing.T) {
	Convey(
		"When creating ctx from a GET request to `/test/:id/:type`",
		t,
		func() {
			testUrl := "/test/:id/:num"
			Convey(
				"And the path params are `id`=`testId` and `num`=`5`",
				func() {
					id := "testId"
					num := 5
					ctx := unittesting.CreateEchoContextFromRequest(
						http.MethodGet,
						testUrl,
						&httptest.ResponseRecorder{},
						unittesting.WithPathParams(map[string]string{
							"id":  id,
							"num": fmt.Sprint(num),
						}),
					)

					Convey("Should have the path `/test/:id/:num`", func() {
						So(ctx.Path(), ShouldEqual, testUrl)
					})

					Convey(
						"Should successfully bind to a valid struct",
						func() {
							dest := struct {
								Id  string `param:"id"`
								Num int    `param:"num"`
							}{}

							err := ctx.Bind(&dest)
							So(err, ShouldBeNil)

							Convey(
								"And the struct should have the correct values",
								func() {
									So(dest.Id, ShouldEqual, id)
									So(dest.Num, ShouldEqual, num)
								},
							)
						},
					)
				},
			)
		},
	)
}

func TestWithQueryParams(t *testing.T) {
	Convey(
		"When creating ctx from a GET request to `/test`",
		t,
		func() {
			testUrl := "/test"
			Convey(
				"And the query params are `id`=`testId` and `num`=`5`",
				func() {
					id := "testId"
					num := 5
					ctx := unittesting.CreateEchoContextFromRequest(
						http.MethodGet,
						testUrl,
						&httptest.ResponseRecorder{},
						unittesting.WithQueryParams(map[string]string{
							"id":  id,
							"num": fmt.Sprint(num),
						}),
					)

					Convey("Should have the path `/test`", func() {
						So(ctx.Path(), ShouldEqual, testUrl)
					})

					Convey(
						"Should successfully bind to a valid struct",
						func() {
							dest := struct {
								Id  string `query:"id"`
								Num int    `query:"num"`
							}{}

							err := ctx.Bind(&dest)
							So(err, ShouldBeNil)

							Convey(
								"And the struct should have the correct values",
								func() {
									So(dest.Id, ShouldEqual, id)
									So(dest.Num, ShouldEqual, num)
								},
							)
						},
					)
				},
			)
		},
	)
}

func TestWithFormPayload(t *testing.T) {
	Convey(
		"When creating ctx from a POST request to `/test` with form data",
		t,
		func() {
			testUrl := "/test"
			Convey(
				"And the form payloads are `id`=`testId` and `num`=`5`",
				func() {
					id := "testId"
					num := 5
					ctx := unittesting.CreateEchoContextFromRequest(
						http.MethodPost,
						testUrl,
						&httptest.ResponseRecorder{},
						unittesting.WithFormPayload(map[string]string{
							"id":  id,
							"num": fmt.Sprint(num),
						}),
					)

					Convey("Should have the path `/test`", func() {
						So(ctx.Path(), ShouldEqual, testUrl)
					})

					Convey(
						"Should successfully bind to a valid struct",
						func() {
							dest := struct {
								Id  string `form:"id"`
								Num int    `form:"num"`
							}{}

							err := ctx.Bind(&dest)
							So(err, ShouldBeNil)

							Convey(
								"And the struct should have the correct values",
								func() {
									So(dest.Id, ShouldEqual, id)
									So(dest.Num, ShouldEqual, num)
								},
							)
						},
					)
				},
			)
		},
	)
}

func TestWithJsonPayload(t *testing.T) {
	Convey(
		"When creating ctx from a POST request to `/test` with json data",
		t,
		func() {
			testUrl := "/test"
			Convey(
				"And the json payloads are `id`=`testId` and `num`=`5`",
				func() {
					id := "testId"
					num := 5
					ctx := unittesting.CreateEchoContextFromRequest(
						http.MethodPost,
						testUrl,
						&httptest.ResponseRecorder{},
						unittesting.WithJsonPayload(map[string]interface{}{
							"id":  id,
							"num": num,
						}),
					)

					Convey("Should have the path `/test`", func() {
						So(ctx.Path(), ShouldEqual, testUrl)
					})

					Convey(
						"Should successfully bind to a valid struct",
						func() {
							dest := struct {
								Id  string `form:"id"`
								Num int    `form:"num"`
							}{}

							err := ctx.Bind(&dest)
							So(err, ShouldBeNil)

							Convey(
								"And the struct should have the correct values",
								func() {
									So(dest.Id, ShouldEqual, id)
									So(dest.Num, ShouldEqual, num)
								},
							)
						},
					)
				},
			)
		},
	)
}
