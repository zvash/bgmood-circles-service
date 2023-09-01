// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package repository

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type AcceptanceType string

const (
	AcceptanceTypeACCEPTALL AcceptanceType = "ACCEPT_ALL"
	AcceptanceTypeASKFIRST  AcceptanceType = "ASK_FIRST"
	AcceptanceTypeREJECTALL AcceptanceType = "REJECT_ALL"
)

func (e *AcceptanceType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = AcceptanceType(s)
	case string:
		*e = AcceptanceType(s)
	default:
		return fmt.Errorf("unsupported scan type for AcceptanceType: %T", src)
	}
	return nil
}

type NullAcceptanceType struct {
	AcceptanceType AcceptanceType `json:"acceptance_type"`
	Valid          bool           `json:"valid"` // Valid is true if AcceptanceType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullAcceptanceType) Scan(value interface{}) error {
	if value == nil {
		ns.AcceptanceType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.AcceptanceType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullAcceptanceType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.AcceptanceType), nil
}

type CircleType string

const (
	CircleTypeROOM CircleType = "ROOM"
	CircleTypeHALL CircleType = "HALL"
)

func (e *CircleType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = CircleType(s)
	case string:
		*e = CircleType(s)
	default:
		return fmt.Errorf("unsupported scan type for CircleType: %T", src)
	}
	return nil
}

type NullCircleType struct {
	CircleType CircleType `json:"circle_type"`
	Valid      bool       `json:"valid"` // Valid is true if CircleType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCircleType) Scan(value interface{}) error {
	if value == nil {
		ns.CircleType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.CircleType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCircleType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.CircleType), nil
}

type MembershipType string

const (
	MembershipTypeOWNER  MembershipType = "OWNER"
	MembershipTypeADMIN  MembershipType = "ADMIN"
	MembershipTypePOSTER MembershipType = "POSTER"
	MembershipTypeVIEWER MembershipType = "VIEWER"
)

func (e *MembershipType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = MembershipType(s)
	case string:
		*e = MembershipType(s)
	default:
		return fmt.Errorf("unsupported scan type for MembershipType: %T", src)
	}
	return nil
}

type NullMembershipType struct {
	MembershipType MembershipType `json:"membership_type"`
	Valid          bool           `json:"valid"` // Valid is true if MembershipType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullMembershipType) Scan(value interface{}) error {
	if value == nil {
		ns.MembershipType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.MembershipType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullMembershipType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.MembershipType), nil
}

type Circle struct {
	ID              uuid.UUID          `json:"id"`
	OwnerID         uuid.UUID          `json:"owner_id"`
	Title           string             `json:"title"`
	Avatar          pgtype.Text        `json:"avatar"`
	Description     pgtype.Text        `json:"description"`
	CircleType      CircleType         `json:"circle_type"`
	IsPrivate       bool               `json:"is_private"`
	IsFeatured      bool               `json:"is_featured"`
	DisplayDuration int32              `json:"display_duration"`
	CreatedAt       time.Time          `json:"created_at"`
	DeletedAt       pgtype.Timestamptz `json:"deleted_at"`
}

type CircleInvitation struct {
	ID       int64       `json:"id"`
	CircleID uuid.UUID   `json:"circle_id"`
	UserID   uuid.UUID   `json:"user_id"`
	Message  pgtype.Text `json:"message"`
}

type CircleJoinRequest struct {
	ID       int64     `json:"id"`
	CircleID uuid.UUID `json:"circle_id"`
	UserID   uuid.UUID `json:"user_id"`
}

type CircleMember struct {
	ID             int64          `json:"id"`
	CircleID       uuid.UUID      `json:"circle_id"`
	MemberID       uuid.UUID      `json:"member_id"`
	MembershipType MembershipType `json:"membership_type"`
	AcceptanceType AcceptanceType `json:"acceptance_type"`
	CreatedAt      time.Time      `json:"created_at"`
}

type CircleTag struct {
	ID       int64     `json:"id"`
	CircleID uuid.UUID `json:"circle_id"`
	TagID    uuid.UUID `json:"tag_id"`
}

type Mood struct {
	ID            uuid.UUID   `json:"id"`
	PosterID      uuid.UUID   `json:"poster_id"`
	CircleID      uuid.UUID   `json:"circle_id"`
	Description   pgtype.Text `json:"description"`
	Image         string      `json:"image"`
	SetBackground bool        `json:"set_background"`
	CreatedAt     time.Time   `json:"created_at"`
}

type MoodReaction struct {
	ID         int64     `json:"id"`
	MoodID     uuid.UUID `json:"mood_id"`
	UserID     uuid.UUID `json:"user_id"`
	ReactionID int64     `json:"reaction_id"`
	CreatedAt  time.Time `json:"created_at"`
}

type Reaction struct {
	ID                 int64       `json:"id"`
	Name               string      `json:"name"`
	Emoji              pgtype.Text `json:"emoji"`
	TextRepresentation pgtype.Text `json:"text_representation"`
}

type Tag struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
