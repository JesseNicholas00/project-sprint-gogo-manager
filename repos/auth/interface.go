package auth

import "context"

type AuthRepository interface {
	CreateUser(ctx context.Context, user User) (User, error)
	FindUserByEmail(ctx context.Context, email string) (User, error)
	UpdateUser(ctx context.Context, user User) (User, error)
}
