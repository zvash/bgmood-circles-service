// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package repository

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	AddMemberToCircle(ctx context.Context, arg AddMemberToCircleParams) (CircleMember, error)
	AskForWPChangeByCircle(ctx context.Context, arg AskForWPChangeByCircleParams) (CircleMember, error)
	CheckIfMemberExists(ctx context.Context, arg CheckIfMemberExistsParams) (bool, error)
	CreateCircle(ctx context.Context, arg CreateCircleParams) (Circle, error)
	CreateJoinRequest(ctx context.Context, arg CreateJoinRequestParams) (CircleJoinRequest, error)
	CreateMood(ctx context.Context, arg CreateMoodParams) (Mood, error)
	DeleteCircle(ctx context.Context, id uuid.UUID) error
	DenyWPChangeAccessToCircle(ctx context.Context, arg DenyWPChangeAccessToCircleParams) (CircleMember, error)
	ExploreCirclesPaginated(ctx context.Context, arg ExploreCirclesPaginatedParams) ([]Circle, error)
	GetCircle(ctx context.Context, id uuid.UUID) (Circle, error)
	GetCircleMoods(ctx context.Context, circleID uuid.UUID) ([]Mood, error)
	GetCircleMoodsPaginated(ctx context.Context, arg GetCircleMoodsPaginatedParams) ([]Mood, error)
	GetJoinRequest(ctx context.Context, id int64) (CircleJoinRequest, error)
	GetUserCircle(ctx context.Context, arg GetUserCircleParams) (Circle, error)
	GrantWPChangeAccessToCircle(ctx context.Context, arg GrantWPChangeAccessToCircleParams) (CircleMember, error)
	InviteToCircle(ctx context.Context, arg InviteToCircleParams) (CircleInvitation, error)
	KickFromCircle(ctx context.Context, arg KickFromCircleParams) error
	LeaveCircle(ctx context.Context, arg LeaveCircleParams) error
	ListInvitations(ctx context.Context, userID uuid.UUID) ([]Circle, error)
	ListJoinedCircles(ctx context.Context, memberID uuid.UUID) ([]Circle, error)
	ListJoinedCirclesPaginated(ctx context.Context, arg ListJoinedCirclesPaginatedParams) ([]Circle, error)
	ListRequestedCircles(ctx context.Context, userID uuid.UUID) ([]Circle, error)
	ReactToMood(ctx context.Context, arg ReactToMoodParams) (MoodReaction, error)
	RemoveAllCircleMembers(ctx context.Context, circleID uuid.UUID) error
	RemoveAllCircleTags(ctx context.Context, circleID uuid.UUID) error
	RemoveAllInvitations(ctx context.Context, circleID uuid.UUID) error
	RemoveAllJoinRequests(ctx context.Context, circleID uuid.UUID) error
	RemoveInvitation(ctx context.Context, arg RemoveInvitationParams) error
	RemoveJoinRequest(ctx context.Context, arg RemoveJoinRequestParams) error
	RemoveReactToMood(ctx context.Context, arg RemoveReactToMoodParams) error
	SetCircleAccessToAdmin(ctx context.Context) (CircleMember, error)
	SetCircleAccessToOwner(ctx context.Context) (CircleMember, error)
	SetCircleAccessToPoster(ctx context.Context) (CircleMember, error)
	SetCircleAccessToViewer(ctx context.Context) (CircleMember, error)
	SetCircleOwner(ctx context.Context, arg SetCircleOwnerParams) (CircleMember, error)
	UpdateCircle(ctx context.Context, arg UpdateCircleParams) (Circle, error)
}

var _ Querier = (*Queries)(nil)
