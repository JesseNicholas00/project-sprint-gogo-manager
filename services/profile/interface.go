package profile

import "context"

type ProfileService interface {
	UpsertUser(
		ctx context.Context,
		req UpsertUserReq,
		res *UpsertUserRes,
	) error
}
