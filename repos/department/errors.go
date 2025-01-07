package department

import "errors"

var (
	ErrDepartmentNotFound = errors.New("departmentRepository: no such department found")
)
