package employee

import "github.com/google/uuid"

type GetEmployeeReq struct {
	Limit          *int   `json:"limit"`
	Offset         *int   `json:"offset"`
	IdentityNumber string `json:"identityNumber"`
	Name           string `json:"name"`
	Gender         string `json:"gender"`
	DepartementId  *int   `json:"departmentId"`
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
	IdentityNumber      string    `json:"identityNumber" validate:"min=5,max=33"`
	Name                string    `json:"name" validate:"min=4,max=33"`
	EmployeeImageUri    string    `json:"employeeImageUri" validate:"url"`
	Gender              string    `json:"gender" validate:"oneof=male female"`
	DepartmentId        uuid.UUID `json:"departmentId"`
	UserID              uuid.UUID `json:"userId"`
	ParamIdentityNumber string
}

type AddEmployeeRes struct {
	IdentityNumber   string    `json:"identityNumber"`
	Name             string    `json:"name"`
	EmployeeImageUri string    `json:"employeeImageUri"`
	Gender           string    `json:"gender"`
	DepartmentId     uuid.UUID `json:"departmentId"`
}
