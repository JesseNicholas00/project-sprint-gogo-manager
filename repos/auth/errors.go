package auth

import "errors"

var ErrEmailNotFound = errors.New(
	"authRepository: no such email found",
)
