package validation

import (
	types "github.com/JesseNicholas00/GogoManager/types/optional"
	"github.com/JesseNicholas00/GogoManager/utils/validation/optional"
	"github.com/JesseNicholas00/GogoManager/utils/validation/uri"
	"reflect"
	"strings"

	"github.com/JesseNicholas00/GogoManager/utils/validation/image"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type EchoValidator struct {
	validator *validator.Validate
}

func (e *EchoValidator) Validate(i interface{}) error {
	return e.validator.Struct(i)
}

var customFields = []customField{
	{
		Tag:       "imageExt",
		Validator: image.ValidateImageExtension,
	},
	{
		Tag:       "complete_uri",
		Validator: uri.ValidateCompleteURI,
	},
}

type customField struct {
	Tag       string
	Validator validator.Func
}

var customTypes = []customType{
	{
		Type:      types.OptionalStr{},
		Validator: optional.ValidateOptionalString,
	},
	{
		Type:      types.OptionalUUID{},
		Validator: optional.ValidateOptionalUUID,
	},
}

type customType struct {
	Type      any
	Validator validator.CustomTypeFunc
}

func NewEchoValidator() echo.Validator {
	validator := validator.New(validator.WithRequiredStructEnabled())

	validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	for _, customField := range customFields {
		validator.RegisterValidation(customField.Tag, customField.Validator)
	}

	for _, customType := range customTypes {
		validator.RegisterCustomTypeFunc(customType.Validator, customType.Type)
	}

	return &EchoValidator{
		validator: validator,
	}
}
