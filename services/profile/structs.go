package profile

type UpsertUserReq struct {
	Email           string `json:"email"    validate:"required,email"`
	Name            string `json:"name"    validate:"required,min=4,max=52"`
	UserImageUri    string `json:"userImageUri"    validate:"required,url"`
	CompanyName     string `json:"companyName"    validate:"required,min=4,max=52"`
	CompanyImageUri string `json:"companyImageUri"    validate:"required,url"`
	Action          string `json:"action" validate:"required,oneof=create login"`
}

type UpsertUserRes struct {
	Email           string `json:"email"`
	Name            string `json:"name"`
	UserImageUri    string `json:"userImageUri"`
	CompanyName     string `json:"companyName"`
	CompanyImageUri string `json:"companyImageUri"`
}
