package uri

import (
	"github.com/go-playground/validator/v10"
	"net/url"
	"strings"
)

func validateCompleteURI(uri string) bool {
	parsed, err := url.ParseRequestURI(uri)
	if err != nil {
		return false
	}

	// Check if the Host seems valid
	return strings.Contains(parsed.Host, ".")
}

func ValidateCompleteURI(fl validator.FieldLevel) bool {
	return validateCompleteURI(fl.Field().String())
}
