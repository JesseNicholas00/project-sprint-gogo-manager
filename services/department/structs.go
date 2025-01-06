package department

import "github.com/google/uuid"

type AddDepartmentReq struct {
	Name string `json:"name" validate:"required,min=4,max=33"`
}

type AddDepartmentRes struct {
	DepartmentId uuid.UUID `json:"departmentId"`
	Name         string    `json:"name"`
}
