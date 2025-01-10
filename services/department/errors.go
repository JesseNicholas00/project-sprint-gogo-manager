package department

import "errors"

var (
	ErrDepartmentNotFound = errors.New("departmentService: no such department found")
	ErrContainEmployee    = errors.New("departmentService: department still has employee(s)")
)
