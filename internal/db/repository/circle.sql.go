// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: circle.sql

package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const addMemberToCircle = `-- name: AddMemberToCircle :one
INSERT INTO circle_members (circle_id, member_id, membership_type, acceptance_type)
VALUES ($1, $2, $3, $4)
RETURNING id, circle_id, member_id, membership_type, acceptance_type, created_at
`

type AddMemberToCircleParams struct {
	CircleID       uuid.UUID      `json:"circle_id"`
	MemberID       uuid.UUID      `json:"member_id"`
	MembershipType MembershipType `json:"membership_type"`
	AcceptanceType AcceptanceType `json:"acceptance_type"`
}

func (q *Queries) AddMemberToCircle(ctx context.Context, arg AddMemberToCircleParams) (CircleMember, error) {
	row := q.db.QueryRow(ctx, addMemberToCircle,
		arg.CircleID,
		arg.MemberID,
		arg.MembershipType,
		arg.AcceptanceType,
	)
	var i CircleMember
	err := row.Scan(
		&i.ID,
		&i.CircleID,
		&i.MemberID,
		&i.MembershipType,
		&i.AcceptanceType,
		&i.CreatedAt,
	)
	return i, err
}

const askForWPChangeByCircle = `-- name: AskForWPChangeByCircle :one
UPDATE circle_members
SET acceptance_type = 'ASK_FIRST'
WHERE circle_id = $1
  AND member_id = $2
RETURNING id, circle_id, member_id, membership_type, acceptance_type, created_at
`

type AskForWPChangeByCircleParams struct {
	CircleID uuid.UUID `json:"circle_id"`
	MemberID uuid.UUID `json:"member_id"`
}

func (q *Queries) AskForWPChangeByCircle(ctx context.Context, arg AskForWPChangeByCircleParams) (CircleMember, error) {
	row := q.db.QueryRow(ctx, askForWPChangeByCircle, arg.CircleID, arg.MemberID)
	var i CircleMember
	err := row.Scan(
		&i.ID,
		&i.CircleID,
		&i.MemberID,
		&i.MembershipType,
		&i.AcceptanceType,
		&i.CreatedAt,
	)
	return i, err
}

const checkIfMemberExists = `-- name: CheckIfMemberExists :one
SELECT EXISTS(SELECT 1 FROM circle_members WHERE circle_id = $1 AND member_id = $2)
`

type CheckIfMemberExistsParams struct {
	CircleID uuid.UUID `json:"circle_id"`
	MemberID uuid.UUID `json:"member_id"`
}

func (q *Queries) CheckIfMemberExists(ctx context.Context, arg CheckIfMemberExistsParams) (bool, error) {
	row := q.db.QueryRow(ctx, checkIfMemberExists, arg.CircleID, arg.MemberID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createCircle = `-- name: CreateCircle :one
INSERT INTO circles (id, owner_id, title, avatar, description, circle_type, is_private, is_featured)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, owner_id, title, avatar, description, circle_type, is_private, is_featured, display_duration, created_at, deleted_at
`

type CreateCircleParams struct {
	ID          uuid.UUID   `json:"id"`
	OwnerID     uuid.UUID   `json:"owner_id"`
	Title       string      `json:"title"`
	Avatar      pgtype.Text `json:"avatar"`
	Description pgtype.Text `json:"description"`
	CircleType  CircleType  `json:"circle_type"`
	IsPrivate   bool        `json:"is_private"`
	IsFeatured  bool        `json:"is_featured"`
}

func (q *Queries) CreateCircle(ctx context.Context, arg CreateCircleParams) (Circle, error) {
	row := q.db.QueryRow(ctx, createCircle,
		arg.ID,
		arg.OwnerID,
		arg.Title,
		arg.Avatar,
		arg.Description,
		arg.CircleType,
		arg.IsPrivate,
		arg.IsFeatured,
	)
	var i Circle
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.Title,
		&i.Avatar,
		&i.Description,
		&i.CircleType,
		&i.IsPrivate,
		&i.IsFeatured,
		&i.DisplayDuration,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const createJoinRequest = `-- name: CreateJoinRequest :one
INSERT INTO circle_join_requests (circle_id, user_id)
VALUES ($1, $2)
RETURNING id, circle_id, user_id
`

type CreateJoinRequestParams struct {
	CircleID uuid.UUID `json:"circle_id"`
	UserID   uuid.UUID `json:"user_id"`
}

func (q *Queries) CreateJoinRequest(ctx context.Context, arg CreateJoinRequestParams) (CircleJoinRequest, error) {
	row := q.db.QueryRow(ctx, createJoinRequest, arg.CircleID, arg.UserID)
	var i CircleJoinRequest
	err := row.Scan(&i.ID, &i.CircleID, &i.UserID)
	return i, err
}

const deleteCircle = `-- name: DeleteCircle :exec
DELETE
FROM circles
WHERE id = $1
`

func (q *Queries) DeleteCircle(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteCircle, id)
	return err
}

const denyWPChangeAccessToCircle = `-- name: DenyWPChangeAccessToCircle :one
UPDATE circle_members
SET acceptance_type = 'REJECT_ALL'
WHERE circle_id = $1
  AND member_id = $2
RETURNING id, circle_id, member_id, membership_type, acceptance_type, created_at
`

type DenyWPChangeAccessToCircleParams struct {
	CircleID uuid.UUID `json:"circle_id"`
	MemberID uuid.UUID `json:"member_id"`
}

func (q *Queries) DenyWPChangeAccessToCircle(ctx context.Context, arg DenyWPChangeAccessToCircleParams) (CircleMember, error) {
	row := q.db.QueryRow(ctx, denyWPChangeAccessToCircle, arg.CircleID, arg.MemberID)
	var i CircleMember
	err := row.Scan(
		&i.ID,
		&i.CircleID,
		&i.MemberID,
		&i.MembershipType,
		&i.AcceptanceType,
		&i.CreatedAt,
	)
	return i, err
}

const exploreCirclesPaginated = `-- name: ExploreCirclesPaginated :many
SELECT id, owner_id, title, avatar, description, circle_type, is_private, is_featured, display_duration, created_at, deleted_at
FROM circles
WHERE circle_type = 'HALL'
ORDER BY created_at DESC
OFFSET $1 LIMIT $2
`

type ExploreCirclesPaginatedParams struct {
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

func (q *Queries) ExploreCirclesPaginated(ctx context.Context, arg ExploreCirclesPaginatedParams) ([]Circle, error) {
	rows, err := q.db.Query(ctx, exploreCirclesPaginated, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Circle{}
	for rows.Next() {
		var i Circle
		if err := rows.Scan(
			&i.ID,
			&i.OwnerID,
			&i.Title,
			&i.Avatar,
			&i.Description,
			&i.CircleType,
			&i.IsPrivate,
			&i.IsFeatured,
			&i.DisplayDuration,
			&i.CreatedAt,
			&i.DeletedAt,
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

const getCircle = `-- name: GetCircle :one
SELECT id, owner_id, title, avatar, description, circle_type, is_private, is_featured, display_duration, created_at, deleted_at
FROM circles
WHERE id = $1
`

func (q *Queries) GetCircle(ctx context.Context, id uuid.UUID) (Circle, error) {
	row := q.db.QueryRow(ctx, getCircle, id)
	var i Circle
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.Title,
		&i.Avatar,
		&i.Description,
		&i.CircleType,
		&i.IsPrivate,
		&i.IsFeatured,
		&i.DisplayDuration,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getJoinRequest = `-- name: GetJoinRequest :one
SELECt id, circle_id, user_id
FROM circle_join_requests
WHERE id = $1
`

func (q *Queries) GetJoinRequest(ctx context.Context, id int64) (CircleJoinRequest, error) {
	row := q.db.QueryRow(ctx, getJoinRequest, id)
	var i CircleJoinRequest
	err := row.Scan(&i.ID, &i.CircleID, &i.UserID)
	return i, err
}

const getUserCircle = `-- name: GetUserCircle :one
SELECT id, owner_id, title, avatar, description, circle_type, is_private, is_featured, display_duration, created_at, deleted_at
FROM circles
WHERE owner_id = $1
  AND id = $2
`

type GetUserCircleParams struct {
	OwnerID uuid.UUID `json:"owner_id"`
	ID      uuid.UUID `json:"id"`
}

func (q *Queries) GetUserCircle(ctx context.Context, arg GetUserCircleParams) (Circle, error) {
	row := q.db.QueryRow(ctx, getUserCircle, arg.OwnerID, arg.ID)
	var i Circle
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.Title,
		&i.Avatar,
		&i.Description,
		&i.CircleType,
		&i.IsPrivate,
		&i.IsFeatured,
		&i.DisplayDuration,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const grantWPChangeAccessToCircle = `-- name: GrantWPChangeAccessToCircle :one
UPDATE circle_members
SET acceptance_type = 'ACCEPT_ALL'
WHERE circle_id = $1
  AND member_id = $2
RETURNING id, circle_id, member_id, membership_type, acceptance_type, created_at
`

type GrantWPChangeAccessToCircleParams struct {
	CircleID uuid.UUID `json:"circle_id"`
	MemberID uuid.UUID `json:"member_id"`
}

func (q *Queries) GrantWPChangeAccessToCircle(ctx context.Context, arg GrantWPChangeAccessToCircleParams) (CircleMember, error) {
	row := q.db.QueryRow(ctx, grantWPChangeAccessToCircle, arg.CircleID, arg.MemberID)
	var i CircleMember
	err := row.Scan(
		&i.ID,
		&i.CircleID,
		&i.MemberID,
		&i.MembershipType,
		&i.AcceptanceType,
		&i.CreatedAt,
	)
	return i, err
}

const inviteToCircle = `-- name: InviteToCircle :one
INSERT INTO circle_invitations (circle_id, user_id)
VALUES ($1, $2)
RETURNING id, circle_id, user_id, message
`

type InviteToCircleParams struct {
	CircleID uuid.UUID `json:"circle_id"`
	UserID   uuid.UUID `json:"user_id"`
}

func (q *Queries) InviteToCircle(ctx context.Context, arg InviteToCircleParams) (CircleInvitation, error) {
	row := q.db.QueryRow(ctx, inviteToCircle, arg.CircleID, arg.UserID)
	var i CircleInvitation
	err := row.Scan(
		&i.ID,
		&i.CircleID,
		&i.UserID,
		&i.Message,
	)
	return i, err
}

const kickFromCircle = `-- name: KickFromCircle :exec
DELETE
FROM circle_members
WHERE circle_id = $1
  AND member_id = $2
`

type KickFromCircleParams struct {
	CircleID uuid.UUID `json:"circle_id"`
	MemberID uuid.UUID `json:"member_id"`
}

func (q *Queries) KickFromCircle(ctx context.Context, arg KickFromCircleParams) error {
	_, err := q.db.Exec(ctx, kickFromCircle, arg.CircleID, arg.MemberID)
	return err
}

const leaveCircle = `-- name: LeaveCircle :exec
DELETE
FROM circle_members
WHERE circle_id = $1
  AND member_id = $2
`

type LeaveCircleParams struct {
	CircleID uuid.UUID `json:"circle_id"`
	MemberID uuid.UUID `json:"member_id"`
}

func (q *Queries) LeaveCircle(ctx context.Context, arg LeaveCircleParams) error {
	_, err := q.db.Exec(ctx, leaveCircle, arg.CircleID, arg.MemberID)
	return err
}

const listInvitations = `-- name: ListInvitations :many
SELECT c.id, c.owner_id, c.title, c.avatar, c.description, c.circle_type, c.is_private, c.is_featured, c.display_duration, c.created_at, c.deleted_at
FROM circles c,
     circle_invitations ci
WHERE c.id = ci.circle_id
  AND ci.user_id = $1
`

func (q *Queries) ListInvitations(ctx context.Context, userID uuid.UUID) ([]Circle, error) {
	rows, err := q.db.Query(ctx, listInvitations, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Circle{}
	for rows.Next() {
		var i Circle
		if err := rows.Scan(
			&i.ID,
			&i.OwnerID,
			&i.Title,
			&i.Avatar,
			&i.Description,
			&i.CircleType,
			&i.IsPrivate,
			&i.IsFeatured,
			&i.DisplayDuration,
			&i.CreatedAt,
			&i.DeletedAt,
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

const listJoinedCircles = `-- name: ListJoinedCircles :many
SELECT c.id, c.owner_id, c.title, c.avatar, c.description, c.circle_type, c.is_private, c.is_featured, c.display_duration, c.created_at, c.deleted_at
FROM circles c,
     circle_members cm
WHERE c.id = cm.circle_id
  AND cm.member_id = $1
`

func (q *Queries) ListJoinedCircles(ctx context.Context, memberID uuid.UUID) ([]Circle, error) {
	rows, err := q.db.Query(ctx, listJoinedCircles, memberID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Circle{}
	for rows.Next() {
		var i Circle
		if err := rows.Scan(
			&i.ID,
			&i.OwnerID,
			&i.Title,
			&i.Avatar,
			&i.Description,
			&i.CircleType,
			&i.IsPrivate,
			&i.IsFeatured,
			&i.DisplayDuration,
			&i.CreatedAt,
			&i.DeletedAt,
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

const listRequestedCircles = `-- name: ListRequestedCircles :many
SELECT c.id, c.owner_id, c.title, c.avatar, c.description, c.circle_type, c.is_private, c.is_featured, c.display_duration, c.created_at, c.deleted_at
FROM circles c,
     circle_join_requests cjr
WHERE cjr.circle_id = c.id
  AND cjr.user_id = $1
`

func (q *Queries) ListRequestedCircles(ctx context.Context, userID uuid.UUID) ([]Circle, error) {
	rows, err := q.db.Query(ctx, listRequestedCircles, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Circle{}
	for rows.Next() {
		var i Circle
		if err := rows.Scan(
			&i.ID,
			&i.OwnerID,
			&i.Title,
			&i.Avatar,
			&i.Description,
			&i.CircleType,
			&i.IsPrivate,
			&i.IsFeatured,
			&i.DisplayDuration,
			&i.CreatedAt,
			&i.DeletedAt,
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

const removeAllCircleMembers = `-- name: RemoveAllCircleMembers :exec
DELETE
FROM circle_members
WHERE circle_id = $1
`

func (q *Queries) RemoveAllCircleMembers(ctx context.Context, circleID uuid.UUID) error {
	_, err := q.db.Exec(ctx, removeAllCircleMembers, circleID)
	return err
}

const removeAllCircleTags = `-- name: RemoveAllCircleTags :exec
DELETE
FROM circle_tag
WHERE circle_id = $1
`

func (q *Queries) RemoveAllCircleTags(ctx context.Context, circleID uuid.UUID) error {
	_, err := q.db.Exec(ctx, removeAllCircleTags, circleID)
	return err
}

const removeAllInvitations = `-- name: RemoveAllInvitations :exec
DELETE
FROM circle_invitations
WHERE circle_id = $1
`

func (q *Queries) RemoveAllInvitations(ctx context.Context, circleID uuid.UUID) error {
	_, err := q.db.Exec(ctx, removeAllInvitations, circleID)
	return err
}

const removeAllJoinRequests = `-- name: RemoveAllJoinRequests :exec
DELETE
FROM circle_join_requests
WHERE circle_id = $1
`

func (q *Queries) RemoveAllJoinRequests(ctx context.Context, circleID uuid.UUID) error {
	_, err := q.db.Exec(ctx, removeAllJoinRequests, circleID)
	return err
}

const removeInvitation = `-- name: RemoveInvitation :exec
DELETE
FROM circle_invitations
WHERE user_id = $1
  AND circle_id = $2
`

type RemoveInvitationParams struct {
	UserID   uuid.UUID `json:"user_id"`
	CircleID uuid.UUID `json:"circle_id"`
}

func (q *Queries) RemoveInvitation(ctx context.Context, arg RemoveInvitationParams) error {
	_, err := q.db.Exec(ctx, removeInvitation, arg.UserID, arg.CircleID)
	return err
}

const removeJoinRequest = `-- name: RemoveJoinRequest :exec
DELETE
FROM circle_join_requests
WHERE user_id = $1
  AND circle_id = $2
`

type RemoveJoinRequestParams struct {
	UserID   uuid.UUID `json:"user_id"`
	CircleID uuid.UUID `json:"circle_id"`
}

func (q *Queries) RemoveJoinRequest(ctx context.Context, arg RemoveJoinRequestParams) error {
	_, err := q.db.Exec(ctx, removeJoinRequest, arg.UserID, arg.CircleID)
	return err
}

const setCircleAccessToAdmin = `-- name: SetCircleAccessToAdmin :one
UPDATE circle_members
SET membership_type = 'ADMIN'
RETURNING id, circle_id, member_id, membership_type, acceptance_type, created_at
`

func (q *Queries) SetCircleAccessToAdmin(ctx context.Context) (CircleMember, error) {
	row := q.db.QueryRow(ctx, setCircleAccessToAdmin)
	var i CircleMember
	err := row.Scan(
		&i.ID,
		&i.CircleID,
		&i.MemberID,
		&i.MembershipType,
		&i.AcceptanceType,
		&i.CreatedAt,
	)
	return i, err
}

const setCircleAccessToOwner = `-- name: SetCircleAccessToOwner :one
UPDATE circle_members
SET membership_type = 'OWNER'
RETURNING id, circle_id, member_id, membership_type, acceptance_type, created_at
`

func (q *Queries) SetCircleAccessToOwner(ctx context.Context) (CircleMember, error) {
	row := q.db.QueryRow(ctx, setCircleAccessToOwner)
	var i CircleMember
	err := row.Scan(
		&i.ID,
		&i.CircleID,
		&i.MemberID,
		&i.MembershipType,
		&i.AcceptanceType,
		&i.CreatedAt,
	)
	return i, err
}

const setCircleAccessToPoster = `-- name: SetCircleAccessToPoster :one
UPDATE circle_members
SET membership_type = 'POSTER'
RETURNING id, circle_id, member_id, membership_type, acceptance_type, created_at
`

func (q *Queries) SetCircleAccessToPoster(ctx context.Context) (CircleMember, error) {
	row := q.db.QueryRow(ctx, setCircleAccessToPoster)
	var i CircleMember
	err := row.Scan(
		&i.ID,
		&i.CircleID,
		&i.MemberID,
		&i.MembershipType,
		&i.AcceptanceType,
		&i.CreatedAt,
	)
	return i, err
}

const setCircleAccessToViewer = `-- name: SetCircleAccessToViewer :one
UPDATE circle_members
SET membership_type = 'VIEWER'
RETURNING id, circle_id, member_id, membership_type, acceptance_type, created_at
`

func (q *Queries) SetCircleAccessToViewer(ctx context.Context) (CircleMember, error) {
	row := q.db.QueryRow(ctx, setCircleAccessToViewer)
	var i CircleMember
	err := row.Scan(
		&i.ID,
		&i.CircleID,
		&i.MemberID,
		&i.MembershipType,
		&i.AcceptanceType,
		&i.CreatedAt,
	)
	return i, err
}

const setCircleOwner = `-- name: SetCircleOwner :one
INSERT INTO circle_members (circle_id, member_id, membership_type, acceptance_type)
VALUES ($1, $2, 'OWNER', 'ASK_FIRST')
RETURNING id, circle_id, member_id, membership_type, acceptance_type, created_at
`

type SetCircleOwnerParams struct {
	CircleID uuid.UUID `json:"circle_id"`
	MemberID uuid.UUID `json:"member_id"`
}

func (q *Queries) SetCircleOwner(ctx context.Context, arg SetCircleOwnerParams) (CircleMember, error) {
	row := q.db.QueryRow(ctx, setCircleOwner, arg.CircleID, arg.MemberID)
	var i CircleMember
	err := row.Scan(
		&i.ID,
		&i.CircleID,
		&i.MemberID,
		&i.MembershipType,
		&i.AcceptanceType,
		&i.CreatedAt,
	)
	return i, err
}

const updateCircle = `-- name: UpdateCircle :one
UPDATE circles
SET title            = COALESCE($1::varchar, title),
    avatar           = COALESCE($2::text, avatar),
    description      = COALESCE($3::text, description),
    is_private       = COALESCE($4::boolean, is_private),
    is_featured      = COALESCE($5::boolean, is_featured),
    display_duration = COALESCE($6::int, display_duration),
    circle_tyoe      = COALESCE($7::circle_type, circle_type)
WHERE id = $8::uuid
RETURNING id, owner_id, title, avatar, description, circle_type, is_private, is_featured, display_duration, created_at, deleted_at
`

type UpdateCircleParams struct {
	Title           pgtype.Text    `json:"title"`
	Avatar          pgtype.Text    `json:"avatar"`
	Description     pgtype.Text    `json:"description"`
	IsPrivate       pgtype.Bool    `json:"is_private"`
	IsFeatured      pgtype.Bool    `json:"is_featured"`
	DisplayDuration pgtype.Int4    `json:"display_duration"`
	CircleType      NullCircleType `json:"circle_type"`
	ID              uuid.UUID      `json:"id"`
}

func (q *Queries) UpdateCircle(ctx context.Context, arg UpdateCircleParams) (Circle, error) {
	row := q.db.QueryRow(ctx, updateCircle,
		arg.Title,
		arg.Avatar,
		arg.Description,
		arg.IsPrivate,
		arg.IsFeatured,
		arg.DisplayDuration,
		arg.CircleType,
		arg.ID,
	)
	var i Circle
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.Title,
		&i.Avatar,
		&i.Description,
		&i.CircleType,
		&i.IsPrivate,
		&i.IsFeatured,
		&i.DisplayDuration,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}