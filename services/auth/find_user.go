package auth

import (
	"context"

	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/JesseNicholas00/GogoManager/utils/transaction"
)

func (svc *authServiceImpl) FindUser(
	ctx context.Context,
	userId string,
	res *FindUserRes,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	ctx, sess, err := svc.dbRizzer.GetOrAppendTx(ctx)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	return transaction.RunWithAutoCommit(&sess, func() error {
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
	})
}
