package employee

import "errors"

var ErrEmployeeNotFound = errors.New(
	"employeeService: no such employee found",
)
