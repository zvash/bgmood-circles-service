package val

type CreateCircleRequest struct {
	Title       string `json:"title" validate:"required,alphanumunicode"`
	Avatar      string `json:"avatar" validate:"required"`
	Description string `json:"description" validate:"omitempty,alphanumunicode"`
	CircleType  string `json:"circle_type"  validate:"required,oneof=ROOM,HALL"`
	IsPrivate   bool   `json:"is_private" validate:"required"`
}
