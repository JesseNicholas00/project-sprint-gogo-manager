package auth

import (
	"context"

	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
)

func (svc *authServiceImpl) FindUser(
	ctx context.Context,
	userId string,
	res *FindUserRes,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	user, err := svc.repo.FindUserByUserId(ctx, userId)

	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	*res = FindUserRes{
		Email:           user.Email,
		Name:            user.Name,
		UserImageUri:    user.UserImageUri,
		CompanyName:     user.CompanyName,
		CompanyImageUri: user.CompanyImageUri,
	}
	return nil
}
