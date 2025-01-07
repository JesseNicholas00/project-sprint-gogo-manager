package auth

import (
	"context"
	"errors"
	"github.com/JesseNicholas00/GogoManager/repos/auth"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/JesseNicholas00/GogoManager/utils/transaction"
)

func (svc *authServiceImpl) LoginUser(
	ctx context.Context,
	req AuthenticateUserReq,
	res *AuthenticateUserRes,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	ctx, sess, err := svc.dbRizzer.GetOrAppendTx(ctx)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	return transaction.RunWithAutoCommit(&sess, func() error {
		user, err := svc.repo.FindUserByEmail(ctx, req.Email)

		if err != nil {
			switch {
			case errors.Is(err, auth.ErrEmailNotFound):
				return ErrUserNotFound
			default:
				return errorutil.AddCurrentContext(err)
			}
		}

		token, err := svc.generateToken(user)
		if err != nil {
			return errorutil.AddCurrentContext(err)
		}

		*res = AuthenticateUserRes{
			Email:       req.Email,
			AccessToken: token,
		}
		return nil
	})
}
