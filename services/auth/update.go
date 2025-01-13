package auth

import (
	"context"
	"errors"
	"github.com/JesseNicholas00/GogoManager/repos/auth"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/JesseNicholas00/GogoManager/utils/transaction"
)

func (svc *authServiceImpl) UpdateUser(
	ctx context.Context,
	userId string,
	req UpdateUserReq,
	res *UpdateUserRes,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	ctx, sess, err := svc.dbRizzer.GetOrAppendTx(ctx)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	return transaction.RunWithAutoCommit(&sess, func() error {
		updatedUser, err := svc.repo.FindUserByUserId(ctx, userId)

		if err != nil {
			return errorutil.AddCurrentContext(err)
		}

		if req.Email.V != nil && *req.Email.V != updatedUser.Email {
			_, err := svc.repo.FindUserByEmail(ctx, *req.Email.V)

			if err == nil {
				return ErrEmailAlreadyRegistered
			}

			if !errors.Is(err, auth.ErrEmailNotFound) {
				// unexpected kind of error
				return errorutil.AddCurrentContext(err)
			}

			updatedUser.Email = *req.Email.V
		}
		if req.Name.V != nil {
			updatedUser.Name = *req.Name.V
		}
		if req.UserImageUri.V != nil {
			updatedUser.UserImageUri = *req.UserImageUri.V
		}
		if req.CompanyName.V != nil {
			updatedUser.CompanyName = *req.CompanyName.V
		}
		if req.CompanyImageUri.V != nil {
			updatedUser.CompanyImageUri = *req.CompanyImageUri.V
		}

		savedUser, err := svc.repo.UpdateUser(ctx, updatedUser)
		if err != nil {
			return errorutil.AddCurrentContext(err)
		}

		*res = UpdateUserRes{
			Email:           savedUser.Email,
			Name:            savedUser.Name,
			UserImageUri:    savedUser.UserImageUri,
			CompanyName:     savedUser.CompanyName,
			CompanyImageUri: savedUser.CompanyImageUri,
		}
		return nil
	})
}
