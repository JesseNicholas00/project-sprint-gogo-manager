package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ConvertToErrList(err validator.ValidationErrors) (errors []string) {
	for _, err := range err {
		var validationTag string

		tag := err.Tag()
		param := err.Param()

		if param != "" {
			validationTag = fmt.Sprintf("%s=%s", tag, param)
		} else {
			validationTag = tag
		}

		curError := fmt.Sprintf(
			"%s: validation failed on %s",
			err.Field(),
			validationTag,
		)
		errors = append(errors, curError)
	}
	return
}
