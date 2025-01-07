package department

import "errors"

var (
	ErrDepartmentNotFound = errors.New("departmentService: no such department found")
)
