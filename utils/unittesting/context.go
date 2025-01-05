package unittesting

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	"github.com/JesseNicholas00/GogoManager/utils/validation"
	"github.com/labstack/echo/v4"
)

func CreateEchoContextFromRequest(
	method string,
	urlPath string,
	recorder *httptest.ResponseRecorder,
	options ...option,
) echo.Context {
	e := echo.New()
	e.Validator = validation.NewEchoValidator()

	for _, option := range options {
		option.urlTr(&urlPath)
	}

	req := httptest.NewRequest(method, urlPath, nil)

	for _, option := range options {
		option.reqTr(req)
	}

	ctx := e.NewContext(req, recorder)
	// remove query params
	ctx.SetPath(strings.Split(urlPath, "?")[0])

	for _, option := range options {
		option.ctxTr(&ctx)
	}

	return ctx
}

/*
# !!Do not confuse with `WithQueryParams`!!

Use this if you need to:

1. Bind to a struct field with the `param` tag

2. Insert `value` into `ctx.Param(key)`

3. Simulate path variables in URLs, i.e. `:id` in `/users/:id`.
*/
func WithPathParams(
	params map[string]string,
) option {
	var keys []string
	var values []string

	for key, value := range params {
		keys = append(keys, key)
		values = append(values, value)
	}

	return option{
		urlTr: defaultUrlTransformer,
		ctxTr: func(ctxPtr *echo.Context) {
			(*ctxPtr).SetParamNames(keys...)
			(*ctxPtr).SetParamValues(values...)
		},
		reqTr: defaultReqTransformer,
	}
}

/*
# !!Do not confuse with `WithPathParams`!!

Use this if you need to:

1. Bind to a struct field with the `query` tag

2. Insert `value` into `ctx.QueryParam(key)`
*/
func WithQueryParams(
	queryParams map[string]string,
) option {
	q := make(url.Values)

	for key, value := range queryParams {
		q.Set(key, value)
	}

	return option{
		urlTr: func(urlPtr *string) {
			*urlPtr += "?" + q.Encode()
		},
		ctxTr: defaultCtxTransformer,
		reqTr: defaultReqTransformer,
	}
}

func WithFormPayload(
	formData map[string]string,
) option {
	f := make(url.Values)

	for key, value := range formData {
		f.Set(key, value)
	}

	return option{
		urlTr: defaultUrlTransformer,
		ctxTr: defaultCtxTransformer,
		reqTr: func(req *http.Request) {
			req.Body = io.NopCloser(
				strings.NewReader(f.Encode()),
			)
			req.ContentLength = -1
			req.Header.Set(
				echo.HeaderContentType,
				echo.MIMEApplicationForm,
			)
		},
	}
}

func WithJsonPayload(
	payload map[string]interface{},
) option {
	json, err := json.Marshal(payload)
	if err != nil {
		panic(
			fmt.Sprintf(
				"failed to marshal json while building test context: %s",
				err,
			),
		)
	}

	return option{
		urlTr: defaultUrlTransformer,
		ctxTr: defaultCtxTransformer,
		reqTr: func(req *http.Request) {
			req.Body = io.NopCloser(
				bytes.NewBuffer(json),
			)
			req.ContentLength = -1
			req.Header.Set(
				echo.HeaderContentType,
				echo.MIMEApplicationJSON,
			)
		},
	}
}

func WithContextData(
	key string,
	value interface{},
) option {
	return option{
		urlTr: defaultUrlTransformer,
		ctxTr: func(ctxPtr *echo.Context) {
			(*ctxPtr).Set(key, value)
		},
		reqTr: defaultReqTransformer,
	}
}

type option struct {
	urlTr urlTransformer
	ctxTr ctxTransformer
	reqTr reqTransformer
}

type urlTransformer func(urlPtr *string)
type ctxTransformer func(ctxPtr *echo.Context)
type reqTransformer func(req *http.Request)

func defaultUrlTransformer(*string)       {}
func defaultCtxTransformer(*echo.Context) {}
func defaultReqTransformer(*http.Request) {}
