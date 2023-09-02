package db

import (
	"context"
	"github.com/google/uuid"
	"github.com/zvash/bgmood-circles-service/internal/db/repository"
)

func (store *SQLStore) RemoveCircleTransaction(ctx context.Context, circleId uuid.UUID) error {
	err := store.executeTransaction(ctx, func(queries *repository.Queries) error {
		if err := queries.RemoveAllInvitations(ctx, circleId); err != nil {
			return err
		}
		if err := queries.RemoveAllJoinRequests(ctx, circleId); err != nil {
			return err
		}
		if err := queries.RemoveAllCircleMembers(ctx, circleId); err != nil {
			return err
		}
		if err := queries.RemoveAllCircleTags(ctx, circleId); err != nil {
			return err
		}
		if err := queries.DeleteCircle(ctx, circleId); err != nil {
			return err
		}
		return nil
	})
	return err
}
