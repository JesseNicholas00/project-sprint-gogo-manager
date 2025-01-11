package employee

import "errors"

var (
	ErrEmployeeNotFound = errors.New("employeeRepository: no such employee found")
	ErrIdentityNumberNotFound = errors.New("employeeRepository: no such identity number found")
)
