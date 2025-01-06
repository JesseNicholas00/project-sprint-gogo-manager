package auth

import (
	"context"
	"errors"

	"github.com/JesseNicholas00/GogoManager/repos/auth"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/JesseNicholas00/GogoManager/utils/transaction"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (svc *authServiceImpl) RegisterUser(
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
		_, err := svc.repo.FindUserByEmail(ctx, req.Email)

		if err == nil {
			return ErrEmailAlreadyRegistered
		}

		if !errors.Is(err, auth.ErrEmailNotFound) {
			// unexpected kind of error
			return errorutil.AddCurrentContext(err)
		}

		cryptedPw, err := bcrypt.GenerateFromPassword(
			[]byte(req.Password),
			svc.bcryptCost,
		)
		if err != nil {
			return errorutil.AddCurrentContext(err)
		}

		savedUser, err := svc.repo.CreateUser(ctx, auth.User{
			Id:       uuid.New().String(),
			Email:    req.Email,
			Password: string(cryptedPw),
		})
		if err != nil {
			return errorutil.AddCurrentContext(err)
		}

		token, err := svc.generateToken(savedUser)
		if err != nil {
			return errorutil.AddCurrentContext(err)
		}

		*res = AuthenticateUserRes{
			Email:       savedUser.Email,
			AccessToken: token,
		}
		return nil
	})
}
