-- name: CreateMood :one
INSERT INTO moods (id, poster_id, circle_id, description, image, set_background)
VALUES (sqlc.arg(id)::uuid,
        sqlc.arg(poster_id)::uuid,
        sqlc.arg(circle_id)::uuid,
        COALESCE(sqlc.narg(description)::text, description),
        sqlc.arg(image)::uuid,
        sqlc.arg(set_background)::boolean)
RETURNING *;

-- name: GetCircleMoods :many
SELECT *
FROM moods
WHERE circle_id = $1
ORDER BY created_at DESC;

-- name: GetCircleMoodsPaginated :many
SELECT *
FROM moods
WHERE circle_id = $1
ORDER BY created_at DESC
OFFSET $2 LIMIT $3;

-- name: ReactToMood :one
INSERT INTO mood_reactions (mood_id, user_id, reaction_id)
VALUES ($1, $2, $3)
ON CONFLICT ("mood_id", "user_id") DO UPDATE SET reaction_id = $3
RETURNING *;

-- name: RemoveReactToMood :exec
DELETE
FROM mood_reactions
WHERE mood_id = $1
  AND user_id = $2;