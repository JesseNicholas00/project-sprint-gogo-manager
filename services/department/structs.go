package department

import "github.com/google/uuid"

type AddDepartmentReq struct {
	Name string `json:"name" validate:"required,min=4,max=33"`
}

type AddDepartmentRes struct {
	DepartmentId uuid.UUID `json:"departmentId"`
	Name         string    `json:"name"`
}

type GetDepartmentParams struct {
	Limit  *uint  `query:"limit"`
	Offset *uint  `query:"offset"`
	Name   string `query:"name"`
}

type GetDepartmentRes []AddDepartmentRes
