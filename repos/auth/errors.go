package auth

import "errors"

var ErrEmailNotFound = errors.New(
	"authRepository: no such email found",
)

var ErrUserIdNotFound = errors.New(
	"authRepository: no such userId found",
)
