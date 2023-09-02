package val

type CreateCircleRequest struct {
	OwnerId     string `json:"owner_id" validate:"required,uuid4"`
	Title       string `json:"title" validate:"required,alphanumunicode"`
	Avatar      string `json:"avatar" validate:"required"`
	Description string `json:"description" validate:"omitempty,alphanumunicode"`
	CircleType  string `json:"circle_type"  validate:"required,oneof=ROOM,HALL"`
	IsPrivate   bool   `json:"is_private" validate:"required"`
}
