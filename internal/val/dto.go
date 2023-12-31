package val

type CreateCircleRequest struct {
	Title       string `json:"title" validate:"required"`
	Avatar      string `json:"avatar" validate:"required"`
	Description string `json:"description" validate:"omitempty"`
	CircleType  string `json:"circle_type"  validate:"required,oneof=ROOM HALL"`
	IsPrivate   bool   `json:"is_private" validate:"omitempty"`
}

type EditCircleRequest struct {
	Title       string `json:"title" validate:"omitempty"`
	Avatar      string `json:"avatar" validate:"omitempty"`
	Description string `json:"description" validate:"omitempty"`
	CircleType  string `json:"circle_type"  validate:"omitempty,oneof=ROOM HALL"`
	IsPrivate   bool   `json:"is_private" validate:"omitempty"`
}

type InviteUserToCircleRequest struct {
	CircleID string `json:"circle_id" validate:"required,uuid"`
	UserID   string `json:"user_id" validate:"required,uuid"`
}

type KickFromCircleRequest struct {
	CircleID string `json:"circle_id" validate:"required,uuid"`
	UserID   string `json:"user_id" validate:"required,uuid"`
}

type SendJoinRequestRequest struct {
	CircleID string `json:"circle_id" validate:"required,uuid"`
}

type AcceptJoinRequestRequest struct {
	JoinRequestID int64 `json:"join_request_id" validate:"required,number,min=1"`
}

type RemoveJoinRequestRequest struct {
	JoinRequestID int64 `json:"join_request_id" validate:"required,number,min=1"`
}

type LeaveCircleRequest struct {
	CircleID string `json:"circle_id" validate:"required,uuid"`
}

type GetJoinedCirclesRequest struct {
	Page int64 `json:"page" validate:"required,number,min=1"`
}

type GetRequestedCirclesRequest struct {
	Page int64 `json:"page" validate:"required,number,min=1"`
}

type ViewCircleRequest struct {
	CircleID string `json:"circle_id" validate:"required,uuid"`
}

type ExploreCirclesRequest struct {
	Page int64 `json:"page" validate:"required,number,min=1"`
}

type CreateMoodRequest struct {
	CircleID      string `json:"circle_id" validate:"required,uuid"`
	Image         string `json:"image" validate:"required"`
	SetBackground bool   `json:"set_background" validate:"omitempty"`
	Description   string `json:"description" validate:"omitempty"`
}

type PromoteToPosterRequest struct {
	CircleID string `json:"circle_id" validate:"required,uuid"`
	UserID   string `json:"user_id" validate:"required,uuid"`
}

type PromoteToAdminRequest struct {
	CircleID string `json:"circle_id" validate:"required,uuid"`
	UserID   string `json:"user_id" validate:"required,uuid"`
}

type DemoteToViewerRequest struct {
	CircleID string `json:"circle_id" validate:"required,uuid"`
	UserID   string `json:"user_id" validate:"required,uuid"`
}

type GetCircleWPChangeAccessRequest struct {
	CircleID string `json:"circle_id" validate:"required,uuid"`
}

type SetCircleWPChangeAccessRequest struct {
	CircleID   string `json:"circle_id" validate:"required,uuid"`
	AccessType string `json:"access_type" validate:"required,oneof=ACCEPT_ALL ASK_FIRST REJECT_ALL"`
}

type ReactRequest struct {
	MoodID     string `json:"mood_id" validate:"required,uuid"`
	ReactionID int64  `json:"reaction_id" validate:"required,number,min=0"`
}

type RemoveReactionRequest struct {
	MoodID string `json:"mood_id" validate:"required,uuid"`
}

type GetCircleJoinRequestsRequest struct {
	CircleID string `json:"circle_id" validate:"required,uuid"`
	Page     int64  `json:"page" validate:"required,number,min=1"`
}
