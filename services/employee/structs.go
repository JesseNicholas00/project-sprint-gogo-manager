package employee

type GetEmployeeReq struct {
	Limit          *int   `json:"limit"`
	Offset         *int   `json:"offset"`
	IdentityNumber string `json:"identityNumber"`
	Name           string `json:"name"`
	Gender         string `json:"gender"`
	DepartementId  *int   `json:"departmentId"`
}

type GetEmployeeResp struct {
	IdentityNumber   string `query:"identityNumber"`
	Name             string `query:"name"`
	EmployeeImageUri string `query:"employeeImageUri"`
	Gender           string `query:"gender"`
	DepartmentId     int    `query:"departmentId"`
}
