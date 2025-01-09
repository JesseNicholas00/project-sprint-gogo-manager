package employee

import "errors"

var ErrIdentityNumberNotFound = errors.New(
	"employeeRepository: no such identity number found",
)
