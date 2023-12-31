// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: mood.sql

package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createMood = `-- name: CreateMood :one
INSERT INTO moods (id, poster_id, circle_id, description, image, set_background)
VALUES ($1::uuid,
        $2::uuid,
        $3::uuid,
        $4::text,
        $5::text,
        $6::boolean)
RETURNING id, poster_id, circle_id, description, image, set_background, created_at
`

type CreateMoodParams struct {
	ID            uuid.UUID   `json:"id"`
	PosterID      uuid.UUID   `json:"poster_id"`
	CircleID      uuid.UUID   `json:"circle_id"`
	Description   pgtype.Text `json:"description"`
	Image         string      `json:"image"`
	SetBackground bool        `json:"set_background"`
}

func (q *Queries) CreateMood(ctx context.Context, arg CreateMoodParams) (Mood, error) {
	row := q.db.QueryRow(ctx, createMood,
		arg.ID,
		arg.PosterID,
		arg.CircleID,
		arg.Description,
		arg.Image,
		arg.SetBackground,
	)
	var i Mood
	err := row.Scan(
		&i.ID,
		&i.PosterID,
		&i.CircleID,
		&i.Description,
		&i.Image,
		&i.SetBackground,
		&i.CreatedAt,
	)
	return i, err
}

const getAvailableReactions = `-- name: GetAvailableReactions :many
SELECT id, name, emoji, text_representation
FROM reactions
ORDER BY id
`

func (q *Queries) GetAvailableReactions(ctx context.Context) ([]Reaction, error) {
	rows, err := q.db.Query(ctx, getAvailableReactions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Reaction{}
	for rows.Next() {
		var i Reaction
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Emoji,
			&i.TextRepresentation,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCircleMoods = `-- name: GetCircleMoods :many
SELECT id, poster_id, circle_id, description, image, set_background, created_at
FROM moods
WHERE circle_id = $1
ORDER BY created_at DESC
`

func (q *Queries) GetCircleMoods(ctx context.Context, circleID uuid.UUID) ([]Mood, error) {
	rows, err := q.db.Query(ctx, getCircleMoods, circleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Mood{}
	for rows.Next() {
		var i Mood
		if err := rows.Scan(
			&i.ID,
			&i.PosterID,
			&i.CircleID,
			&i.Description,
			&i.Image,
			&i.SetBackground,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCircleMoodsPaginated = `-- name: GetCircleMoodsPaginated :many
SELECT id, poster_id, circle_id, description, image, set_background, created_at
FROM moods
WHERE circle_id = $1
ORDER BY created_at DESC
OFFSET $2 LIMIT $3
`

type GetCircleMoodsPaginatedParams struct {
	CircleID uuid.UUID `json:"circle_id"`
	Offset   int32     `json:"offset"`
	Limit    int32     `json:"limit"`
}

func (q *Queries) GetCircleMoodsPaginated(ctx context.Context, arg GetCircleMoodsPaginatedParams) ([]Mood, error) {
	rows, err := q.db.Query(ctx, getCircleMoodsPaginated, arg.CircleID, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Mood{}
	for rows.Next() {
		var i Mood
		if err := rows.Scan(
			&i.ID,
			&i.PosterID,
			&i.CircleID,
			&i.Description,
			&i.Image,
			&i.SetBackground,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMood = `-- name: GetMood :one
SELECT id, poster_id, circle_id, description, image, set_background, created_at
FROM moods
WHERE id = $1
`

func (q *Queries) GetMood(ctx context.Context, id uuid.UUID) (Mood, error) {
	row := q.db.QueryRow(ctx, getMood, id)
	var i Mood
	err := row.Scan(
		&i.ID,
		&i.PosterID,
		&i.CircleID,
		&i.Description,
		&i.Image,
		&i.SetBackground,
		&i.CreatedAt,
	)
	return i, err
}

const getMoodReactions = `-- name: GetMoodReactions :many
SELECT r.id, r.name, r.emoji, r.text_representation, COUNT(mr.reaction_id) as reaction_count
from reactions r
         LEFT JOIN mood_reactions mr ON r.id = mr.reaction_id
WHERE mr.mood_id = $1
GROUP BY r.id
ORDER BY r.id
`

type GetMoodReactionsRow struct {
	ID                 int64       `json:"id"`
	Name               string      `json:"name"`
	Emoji              pgtype.Text `json:"emoji"`
	TextRepresentation pgtype.Text `json:"text_representation"`
	ReactionCount      int64       `json:"reaction_count"`
}

func (q *Queries) GetMoodReactions(ctx context.Context, moodID uuid.UUID) ([]GetMoodReactionsRow, error) {
	rows, err := q.db.Query(ctx, getMoodReactions, moodID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetMoodReactionsRow{}
	for rows.Next() {
		var i GetMoodReactionsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Emoji,
			&i.TextRepresentation,
			&i.ReactionCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserMoodReaction = `-- name: GetUserMoodReaction :one
SELECT id, name, emoji, text_representation
FROM reactions r
WHERE id = (SELECT reaction_id FROM mood_reactions WHERE mood_id = $1 AND user_id = $2 LIMIT 1)
`

type GetUserMoodReactionParams struct {
	MoodID uuid.UUID `json:"mood_id"`
	UserID uuid.UUID `json:"user_id"`
}

func (q *Queries) GetUserMoodReaction(ctx context.Context, arg GetUserMoodReactionParams) (Reaction, error) {
	row := q.db.QueryRow(ctx, getUserMoodReaction, arg.MoodID, arg.UserID)
	var i Reaction
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Emoji,
		&i.TextRepresentation,
	)
	return i, err
}

const reactToMood = `-- name: ReactToMood :one
INSERT INTO mood_reactions (mood_id, user_id, reaction_id)
VALUES ($1, $2, $3)
ON CONFLICT ("mood_id", "user_id") DO UPDATE SET reaction_id = $3
RETURNING id, mood_id, user_id, reaction_id, created_at
`

type ReactToMoodParams struct {
	MoodID     uuid.UUID `json:"mood_id"`
	UserID     uuid.UUID `json:"user_id"`
	ReactionID int64     `json:"reaction_id"`
}

func (q *Queries) ReactToMood(ctx context.Context, arg ReactToMoodParams) (MoodReaction, error) {
	row := q.db.QueryRow(ctx, reactToMood, arg.MoodID, arg.UserID, arg.ReactionID)
	var i MoodReaction
	err := row.Scan(
		&i.ID,
		&i.MoodID,
		&i.UserID,
		&i.ReactionID,
		&i.CreatedAt,
	)
	return i, err
}

const removeReactToMood = `-- name: RemoveReactToMood :exec
DELETE
FROM mood_reactions
WHERE mood_id = $1
  AND user_id = $2
`

type RemoveReactToMoodParams struct {
	MoodID uuid.UUID `json:"mood_id"`
	UserID uuid.UUID `json:"user_id"`
}

func (q *Queries) RemoveReactToMood(ctx context.Context, arg RemoveReactToMoodParams) error {
	_, err := q.db.Exec(ctx, removeReactToMood, arg.MoodID, arg.UserID)
	return err
}
