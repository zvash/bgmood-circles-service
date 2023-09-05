package db

import (
	"context"
	"github.com/zvash/bgmood-circles-service/internal/db/repository"
)

type CreateMoodTransactionParams struct {
	repository.CreateMoodParams
	AfterMoodCreated func(mood repository.Mood) error
}

func (store *SQLStore) CreateMoodTransaction(
	ctx context.Context,
	arg CreateMoodTransactionParams,
) (repository.Mood, error) {
	var mood repository.Mood
	err := store.executeTransaction(ctx, func(queries *repository.Queries) error {
		mood, err := queries.CreateMood(ctx, arg.CreateMoodParams)
		if err != nil {
			return err
		}
		return arg.AfterMoodCreated(mood)
	})
	if err != nil {
		return mood, err
	}
	return mood, nil
}
