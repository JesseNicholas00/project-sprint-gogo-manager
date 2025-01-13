package auth

import (
	"github.com/JesseNicholas00/GogoManager/types/optional"
	"github.com/golang-jwt/jwt/v4"
)

type AuthenticateUserReq struct {
	Password string `json:"password" validate:"required,min=8,max=32"`
	Email    string `json:"email"    validate:"required,email"`
	Action   string `json:"action" validate:"required,oneof=create login"`
}

type AuthenticateUserRes struct {
	Email       string `json:"email"`
	AccessToken string `json:"token"`
}

type UpdateUserReq struct {
	Email           optional.OptionalStr `json:"email" validate:"omitnil,email"`
	Name            optional.OptionalStr `json:"name" validate:"omitnil,min=4,max=52"`
	UserImageUri    optional.OptionalStr `json:"userImageUri" validate:"omitnil,url"`
	CompanyName     optional.OptionalStr `json:"companyName" validate:"omitnil,min=4,max=52"`
	CompanyImageUri optional.OptionalStr `json:"companyImageUri" validate:"omitnil,url"`
}

type UpdateUserRes struct {
	Email           string `json:"email"`
	Name            string `json:"name"`
	UserImageUri    string `json:"userImageUri"`
	CompanyName     string `json:"companyName"`
	CompanyImageUri string `json:"companyImageUri"`
}

type FindUserRes struct {
	Email           string `json:"email"`
	Name            string `json:"name"`
	UserImageUri    string `json:"userImageUri"`
	CompanyName     string `json:"companyName"`
	CompanyImageUri string `json:"companyImageUri"`
}

type GetSessionFromTokenReq struct {
	AccessToken string
}

type GetSessionFromTokenRes struct {
	UserId string
}

type jwtSubClaims struct {
	UserId string `json:"userId"`
}

type jwtClaims struct {
	jwt.RegisteredClaims
	Data jwtSubClaims `json:"data"`
}
