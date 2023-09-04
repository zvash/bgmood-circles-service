package val

type CreateCircleRequest struct {
	Title       string `json:"title" validate:"required,alphanumunicode"`
	Avatar      string `json:"avatar" validate:"required"`
	Description string `json:"description" validate:"omitempty,alphanumunicode"`
	CircleType  string `json:"circle_type"  validate:"required,oneof=ROOM,HALL"`
	IsPrivate   bool   `json:"is_private" validate:"required"`
}

type EditCircleRequest struct {
	Title       string `json:"title" validate:"omitempty,alphanumunicode"`
	Avatar      string `json:"avatar" validate:"omitempty"`
	Description string `json:"description" validate:"omitempty,alphanumunicode"`
	CircleType  string `json:"circle_type"  validate:"omitempty,oneof=ROOM,HALL"`
	IsPrivate   bool   `json:"is_private" validate:"omitempty"`
}

type InviteUserToCircleRequest struct {
	CircleID string `json:"circle_id" validate:"required,uuid"`
	UserID   string `json:"user_id" validate:"required,uuid"`
}
