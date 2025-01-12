package auth

import (
	"github.com/golang-jwt/jwt/v4"
)

type AuthenticateUserReq struct {
	Password string `json:"password" validate:"required,min=5,max=30"`
	Email    string `json:"email"    validate:"required,email"`
	Action   string `json:"action" validate:"required,oneof=create login"`
}

type AuthenticateUserRes struct {
	Email       string `json:"email"`
	AccessToken string `json:"token"`
}

type UpdateUserReq struct {
	Email           *string `json:"email" validate:"omitnil,email"`
	Name            *string `json:"name" validate:"omitnil,min=4,max=52"`
	UserImageUri    *string `json:"userImageUri" validate:"omitnil,url"`
	CompanyName     *string `json:"companyName" validate:"omitnil,min=4,max=52"`
	CompanyImageUri *string `json:"companyImageUri" validate:"omitnil,url"`
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
