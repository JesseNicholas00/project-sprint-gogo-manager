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

		if req.Email != nil {
			_, err := svc.repo.FindUserByEmail(ctx, *req.Email)

			if err == nil {
				return ErrEmailAlreadyRegistered
			}

			if !errors.Is(err, auth.ErrEmailNotFound) {
				// unexpected kind of error
				return errorutil.AddCurrentContext(err)
			}

			updatedUser.Email = *req.Email
		}
		if req.Name != nil {
			updatedUser.Name = *req.Name
		}
		if req.UserImageUri != nil {
			updatedUser.UserImageUri = *req.UserImageUri
		}
		if req.CompanyName != nil {
			updatedUser.CompanyName = *req.CompanyName
		}
		if req.CompanyImageUri != nil {
			updatedUser.CompanyImageUri = *req.CompanyImageUri
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
