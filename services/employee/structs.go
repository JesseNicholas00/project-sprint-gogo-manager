package employee

import "github.com/google/uuid"

type GetEmployeeReq struct {
	Limit          *int   `json:"limit" query:"limit"`
	Offset         *int   `json:"offset" query:"offset"`
	IdentityNumber string `json:"identityNumber" query:"identityNumber"`
	Name           string `json:"name" query:"name"`
	Gender         string `json:"gender" query:"gender"`
	DepartementId  string `json:"departmentId" query:"departmentId"`
}

type GetEmployeeResp []AddEmployeeRes

type AddEmployeeReq struct {
	IdentityNumber   string    `json:"identityNumber" validate:"required,min=5,max=33"`
	Name             string    `json:"name" validate:"required,min=4,max=33"`
	EmployeeImageUri string    `json:"employeeImageUri" validate:"required,url"`
	Gender           string    `json:"gender" validate:"required,oneof=male female"`
	DepartmentId     uuid.UUID `json:"departmentId" validate:"required"`
	UserID           uuid.UUID `json:"userId"`
}

type UpdateEmployeeReq struct {
	IdentityNumber      *string    `json:"identityNumber" validate:"omitnil,min=5,max=33"`
	Name                *string    `json:"name" validate:"omitnil,min=4,max=33"`
	EmployeeImageUri    *string    `json:"employeeImageUri" validate:"omitnil,url"`
	Gender              *string    `json:"gender" validate:"omitnil,oneof=male female"`
	DepartmentId        *uuid.UUID `json:"departmentId"`
	UserID              string     `json:"userId"`
	ParamIdentityNumber string
}

type AddEmployeeRes struct {
	IdentityNumber   string    `json:"identityNumber"`
	Name             string    `json:"name"`
	EmployeeImageUri string    `json:"employeeImageUri"`
	Gender           string    `json:"gender"`
	DepartmentId     uuid.UUID `json:"departmentId"`
}

type DeleteEmployeeReq struct {
	IdentityNumber string `param:"identityNumber"`
}
