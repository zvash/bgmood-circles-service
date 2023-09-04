package db

import (
	"context"
	"github.com/zvash/bgmood-circles-service/internal/db/repository"
)

type InviteToCircleTransactionParams struct {
	repository.InviteToCircleParams
	AfterInvite func(invitation repository.CircleInvitation) error
}

func (store *SQLStore) InviteToCircleTransaction(
	ctx context.Context,
	arg InviteToCircleTransactionParams,
) (repository.CircleInvitation, error) {
	var result repository.CircleInvitation
	err := store.executeTransaction(ctx, func(queries *repository.Queries) error {
		result, err := queries.InviteToCircle(ctx, arg.InviteToCircleParams)
		if err != nil {
			return err
		}
		return arg.AfterInvite(result)
	})
	return result, err
}
