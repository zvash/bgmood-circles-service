-- name: CreateCircle :one
INSERT INTO circles (id, owner_id, title, avatar, description, circle_type, is_private, is_featured)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: DeleteCircle :exec
DELETE
FROM circles
WHERE id = $1;

-- name: GetUserCircle :one
SELECT *
FROM circles
WHERE owner_id = $1
  AND id = $2;

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

-- name: GetJoinRequest :one
SELECt *
FROM circle_join_requests
WHERE id = $1;

-- name: RemoveInvitation :exec
DELETE
FROM circle_invitations
WHERE user_id = $1
  AND circle_id = $2;

-- name: RemoveAllJoinRequests :exec
DELETE
FROM circle_join_requests
WHERE circle_id = $1;

-- name: RemoveAllInvitations :exec
DELETE
FROM circle_invitations
WHERE circle_id = $1;

-- name: AddMemberToCircle :one
INSERT INTO circle_members (circle_id, member_id, membership_type, acceptance_type)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: SetCircleOwner :one
INSERT INTO circle_members (circle_id, member_id, membership_type, acceptance_type)
VALUES ($1, $2, 'OWNER', 'ASK_FIRST')
RETURNING *;

-- name: KickFromCircle :exec
DELETE
FROM circle_members
WHERE circle_id = $1
  AND member_id = $2;

-- name: CheckIfMemberExists :one
SELECT EXISTS(SELECT 1 FROM circle_members WHERE circle_id = $1 AND member_id = $2);

-- name: CheckIfMemberCanPostToCircle :one
SELECT EXISTS(SELECT TRUE
              FROM circle_members
              WHERE circle_id = $1
                AND member_id = $2
                AND membership_type IN ('OWNER', 'POSTER', 'ADMIN'));

-- name: CheckIfMemberCanChangeUsersAccess :one
SELECT EXISTS(SELECT TRUE
              FROM circle_members
              WHERE circle_id = $1
                AND member_id = $2
                AND membership_type IN ('OWNER', 'ADMIN'));

-- name: ListRequestedCircles :many
SELECT c.*
FROM circles c,
     circle_join_requests cjr
WHERE cjr.circle_id = c.id
  AND cjr.user_id = $1;

-- name: ListRequestedCirclesPaginated :many
SELECT c.*
FROM circles c,
     circle_join_requests cjr
WHERE cjr.circle_id = c.id
  AND cjr.user_id = $1
ORDER BY cjr.id DESC
OFFSET $2 LIMIT $3;;

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
  AND cm.member_id = $1
ORDER BY cm.created_at DESC;

-- name: ListJoinedCirclesPaginated :many
SELECT c.*
FROM circles c,
     circle_members cm
WHERE c.id = cm.circle_id
  AND cm.member_id = $1
ORDER BY cm.created_at DESC
OFFSET $2 LIMIT $3;

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
    display_duration = COALESCE(sqlc.narg(display_duration)::int, display_duration),
    circle_tyoe      = COALESCE(sqlc.narg(circle_type)::circle_type, circle_type)
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
WHERE circle_id = $1
  AND member_id = $2
RETURNING *;

-- name: SetCircleAccessToPoster :one
UPDATE circle_members
SET membership_type = 'POSTER'
WHERE circle_id = $1
  AND member_id = $2
RETURNING *;

-- name: SetCircleAccessToViewer :one
UPDATE circle_members
SET membership_type = 'VIEWER'
WHERE circle_id = $1
  AND member_id = $2
RETURNING *;

-- name: SetCircleAccessToOwner :one
UPDATE circle_members
SET membership_type = 'OWNER'
WHERE circle_id = $1
  AND member_id = $2
RETURNING *;

-- name: ExploreCirclesPaginated :many
SELECT *
FROM circles
WHERE circle_type = 'HALL'
ORDER BY created_at DESC
OFFSET $1 LIMIT $2;

-- name: RemoveAllCircleMembers :exec
DELETE
FROM circle_members
WHERE circle_id = $1;

-- name: GetCircleMember :one
SELECT *
FROM circle_members
WHERE circle_id = $1
  AND member_id = $2;

-- name: SetCircleMemberAccess :one
UPDATE circle_members
SET acceptance_type = $3
WHERE circle_id = $1
  AND member_id = $2
RETURNING *;

-- name: RemoveAllCircleTags :exec
DELETE
FROM circle_tag
WHERE circle_id = $1;

-- name: DisplayCircleForUser :one
SELECT c.*, (cm.member_id IS NOT NULL AND cm.member_id = $2) as is_member
FROM circles c
         LEFT JOIN circle_members cm ON c.id = cm.circle_id
WHERE c.id = $1
  AND (cm.member_id = $2 OR c.circle_type = 'HALL')
ORDER BY is_member DESC
LIMIT 1;

-- name: ExploreCirclesForUserPaginated :many
SELECT c.*,
       count(m.id)                                             as mood_count,
       count(cm.circle_id)                                     as member_count,
       bool_or(cm.member_id IS NOT NULL AND cm.member_id = $1) as is_member
FROM circles c
         LEFT JOIN circle_members cm ON c.id = cm.circle_id
         LEFT JOIN moods m ON c.id = m.circle_id
WHERE c.circle_type = 'HALL'
   OR cm.member_id = $1
GROUP BY c.id
ORDER BY c.created_at DESC
OFFSET $2 LIMIT $3;


-- name: GetMemberCount :one
SELECT count(*)
FROM circle_members
WHERE circle_id = $1;

-- name: GetMoodCount :one
SELECT count(*)
FROM moods
WHERE circle_id = $1;