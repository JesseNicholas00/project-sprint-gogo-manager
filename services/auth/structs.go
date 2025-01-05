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
