package image

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var blank = struct{}{}

var imageExtensions = map[string]struct{}{
	"png":  blank,
	"jpg":  blank,
	"jpeg": blank,
}

func validateImageExtensionImpl(imageUrl string) bool {
	ext := imageUrl[strings.LastIndex(imageUrl, ".")+1:]
	_, ok := imageExtensions[ext]
	return ok
}

func ValidateImageExtension(fl validator.FieldLevel) bool {
	return validateImageExtensionImpl(fl.Field().String())
}
