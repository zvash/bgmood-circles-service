package db

import (
	"context"
	"github.com/google/uuid"
	"github.com/zvash/bgmood-circles-service/internal/db/repository"
)

type AcceptJoinRequestTransactionParams struct {
	CircleID    uuid.UUID
	MemberID    uuid.UUID
	AfterAccept func(repository.CircleMember) error
}

func (store *SQLStore) AcceptJoinRequestTransaction(
	ctx context.Context,
	arg AcceptJoinRequestTransactionParams,
) (repository.CircleMember, error) {
	var result repository.CircleMember
	err := store.executeTransaction(ctx, func(queries *repository.Queries) error {

		result, err := queries.AddMemberToCircle(ctx, repository.AddMemberToCircleParams{
			CircleID:       arg.CircleID,
			MemberID:       arg.MemberID,
			MembershipType: repository.MembershipTypeVIEWER,
			AcceptanceType: repository.AcceptanceTypeASKFIRST,
		})
		if err != nil {
			return err
		}
		err = queries.RemoveJoinRequest(ctx, repository.RemoveJoinRequestParams{
			CircleID: arg.CircleID,
			UserID:   arg.MemberID,
		})
		if err != nil {
			return err
		}
		return arg.AfterAccept(result)
	})
	return result, err
}
