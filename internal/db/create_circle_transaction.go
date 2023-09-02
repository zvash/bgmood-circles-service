package db

import (
	"context"
	"github.com/zvash/bgmood-circles-service/internal/db/repository"
)

func (store *SQLStore) CreateCircleTransaction(ctx context.Context, params repository.CreateCircleParams) (repository.Circle, error) {
	var circle repository.Circle
	err := store.executeTransaction(ctx, func(queries *repository.Queries) error {
		var err error
		circle, err = queries.CreateCircle(ctx, params)
		if err != nil {
			return err
		}
		_, err = queries.SetCircleOwner(ctx, repository.SetCircleOwnerParams{
			CircleID: circle.ID,
			MemberID: circle.OwnerID,
		})
		if err != nil {
			return err
		}
		return nil
	})
	return circle, err
}
