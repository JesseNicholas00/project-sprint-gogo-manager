package employee

import "errors"

var (
	ErrEmployeeNotFound            = errors.New("employeeService: no such employee found")
	ErrIdentityNumberNotFound      = errors.New("employeeService: no such identity number found")
	ErrIdentityNumberAlreadyExists = errors.New("employeeService: identity number already exists")
)
