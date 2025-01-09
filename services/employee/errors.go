package employee

import "errors"

var (
	ErrEmployeeNotFound = errors.New("employeeService: no such employee found")
	ErrorEmployeeExist  = errors.New("employeeService: employee already exist")
)
