package employee

import "errors"

var ErrIdentityNumberNotFound = errors.New(
	"employeeRepository: no such identity number found",
)

var ErrIdentityNumberAlreadyExists = errors.New(
	"employeeRepository: identity number already exists",
)
