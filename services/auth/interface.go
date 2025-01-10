package auth

import "context"

type AuthService interface {
	RegisterUser(
		ctx context.Context,
		req AuthenticateUserReq,
		res *AuthenticateUserRes,
	) error
	LoginUser(
		ctx context.Context,
		req AuthenticateUserReq,
		res *AuthenticateUserRes,
	) error
	GetSessionFromToken(
		ctx context.Context,
		req GetSessionFromTokenReq,
		res *GetSessionFromTokenRes,
	) error
	UpdateUser(
		ctx context.Context,
		userId string,
		req UpdateUserReq,
		res *UpdateUserRes,
	) error
}
