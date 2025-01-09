package employee

import "errors"

var (
	ErrEmployeeNotFound = errors.New("employeeRepository: no such employee found")
	ErrorEmployeeExist  = errors.New("employeeRepository: employee already exist")
)
