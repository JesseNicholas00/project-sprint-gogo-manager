package profile

import "errors"

var ErrEmailAlreadyRegistered = errors.New(
	"authService: email already registered",
)
var ErrUserNotFound = errors.New(
	"authService: no such user found",
)
var ErrInvalidCredentials = errors.New(
	"authService: invalid credentials",
)
var ErrTokenInvalid = errors.New(
	"authService: invalid access token",
)
var ErrTokenExpired = errors.New(
	"authService: token expired",
)
