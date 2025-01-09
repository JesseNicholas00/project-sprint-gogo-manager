package employee

import "errors"

var (
	ErrIdentityNumberNotFound      = errors.New("employeeService: no such identity number found")
	ErrIdentityNumberAlreadyExists = errors.New("employeeService: identity number already exists")
)
