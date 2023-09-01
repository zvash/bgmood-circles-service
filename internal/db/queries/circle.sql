-- name: CreateCircle :one
INSERT INTO circles (id, owner_id, title, avatar, description, circle_type, is_private, is_featured)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: DeleteCircle :exec
DELETE
FROM circles
WHERE id = $1;

-- name: InviteToCircle :one
INSERT INTO circle_invitations (circle_id, user_id)
VALUES ($1, $2)
RETURNING *;

-- name: CreateJoinRequest :one
INSERT INTO circle_join_requests (circle_id, user_id)
VALUES ($1, $2)
RETURNING *;

-- name: RemoveJoinRequest :exec
DELETE
FROM circle_join_requests
WHERE user_id = $1
  AND circle_id = $2;

-- name: RemoveInvitation :exec
DELETE
FROM circle_invitations
WHERE user_id = $1
  AND circle_id = $2;

-- name: AcceptInvitation :one
INSERT INTO circle_members (circle_id, member_id, membership_type, acceptance_type)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: KickFromCircle :exec
DELETE
FROM circle_members
WHERE circle_id = $1
  AND member_id = $2;

-- name: ListRequestedCircles :many
SELECT c.*
FROM circles c,
     circle_join_requests cjr
WHERE cjr.circle_id = c.id
  AND cjr.user_id = $1;

-- name: ListInvitations :many
SELECT c.*
FROM circles c,
     circle_invitations ci
WHERE c.id = ci.circle_id
  AND ci.user_id = $1;

-- name: ListJoinedCircles :many
SELECT c.*
FROM circles c,
     circle_members cm
WHERE c.id = cm.circle_id
  AND cm.member_id = $1;

-- name: LeaveCircle :exec
DELETE
FROM circle_members
WHERE circle_id = $1
  AND member_id = $2;

-- name: GetCircle :one
SELECT *
FROM circles
WHERE id = $1;


-- name: UpdateCircle :one
UPDATE circles
SET title            = COALESCE(sqlc.narg(title)::varchar, title),
    avatar           = COALESCE(sqlc.narg(avatar)::text, avatar),
    description      = COALESCE(sqlc.narg(description)::text, description),
    is_private       = COALESCE(sqlc.narg(is_private)::boolean, is_private),
    is_featured      = COALESCE(sqlc.narg(is_featured)::boolean, is_featured),
    display_duration = COALESCE(sqlc.narg(display_duration)::int, display_duration)
WHERE id = sqlc.arg(id)::uuid
RETURNING *;

-- name: GrantWPChangeAccessToCircle :one
UPDATE circle_members
SET acceptance_type = 'ACCEPT_ALL'
WHERE circle_id = $1
  AND member_id = $2
RETURNING *;

-- name: AskForWPChangeByCircle :one
UPDATE circle_members
SET acceptance_type = 'ASK_FIRST'
WHERE circle_id = $1
  AND member_id = $2
RETURNING *;

-- name: DenyWPChangeAccessToCircle :one
UPDATE circle_members
SET acceptance_type = 'REJECT_ALL'
WHERE circle_id = $1
  AND member_id = $2
RETURNING *;

-- name: SetCircleAccessToAdmin :one
UPDATE circle_members
SET membership_type = 'ADMIN'
RETURNING *;

-- name: SetCircleAccessToPoster :one
UPDATE circle_members
SET membership_type = 'POSTER'
RETURNING *;

-- name: SetCircleAccessToViewer :one
UPDATE circle_members
SET membership_type = 'VIEWER'
RETURNING *;

-- name: ExploreCirclesPaginated :many
SELECT *
FROM circles
WHERE circle_type = 'HALL'
ORDER BY created_at DESC
OFFSET $1 LIMIT $2;