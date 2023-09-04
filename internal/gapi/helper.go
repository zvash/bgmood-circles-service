package gapi

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/zvash/bgmood-circles-service/internal/cpb"
	"github.com/zvash/bgmood-circles-service/internal/db"
	"github.com/zvash/bgmood-circles-service/internal/db/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func stringToUUID(value string) (uuid.UUID, error) {
	res, err := uuid.Parse(value)
	if err != nil {
		return res, status.Errorf(codes.InvalidArgument, "invalid uuid: %v", err)
	}
	return res, nil
}

func (server *Server) getOwnerUUID(ctx context.Context) (uuid.UUID, error) {
	var ownerUUID uuid.UUID
	owner, err := server.extractUserFromContext(ctx)
	if err != nil {
		return ownerUUID, status.Errorf(codes.Unauthenticated, "unauthorized.")
	}
	return stringToUUID(owner.Id)
}

func (server *Server) getCircleForOwner(ctx context.Context, circleID uuid.UUID) (repository.Circle, error) {
	var circle repository.Circle
	ownerUUID, err := server.getOwnerUUID(ctx)
	if err != nil {
		return circle, err
	}
	circle, err = server.db.GetUserCircle(ctx, repository.GetUserCircleParams{
		ID:      circleID,
		OwnerID: ownerUUID,
	})
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return circle, status.Errorf(codes.NotFound, "could not find the circle")
		}
		return circle, status.Errorf(codes.Internal, "internal server error.")
	}
	return circle, nil
}

func (server *Server) getCircleMeta(ctx context.Context, circleID uuid.UUID, userID uuid.UUID) (*cpb.CircleMeta, error) {
	circle, err := server.db.DisplayCircleToUser(ctx, repository.DisplayCircleToUserParams{
		ID:       circleID,
		MemberID: userID,
	})
	if err != nil {
		return nil, notFoundOrInternalError(err)
	}
	circleMeta := &cpb.CircleMeta{
		Id:             circle.ID.String(),
		OwnerId:        circle.OwnerID.String(),
		Title:          circle.Title,
		Avatar:         circle.Avatar.String,
		Description:    circle.Description.String,
		CircleType:     string(circle.CircleType),
		IsPrivate:      circle.IsPrivate,
		IsFeatured:     circle.IsFeatured,
		CreatedAt:      timestamppb.New(circle.CreatedAt),
		MoodCount:      0,
		MemberCount:    0,
		ViewerIsMember: circle.IsMember,
	}
	memberCount, err := server.db.GetMemberCount(ctx, circle.ID)
	if err != nil {
		return nil, internalServerError()
	}
	circleMeta.MemberCount = memberCount
	moodCount, err := server.db.GetMoodCount(ctx, circle.ID)
	if err != nil {
		return nil, internalServerError()
	}
	circleMeta.MoodCount = moodCount
	return circleMeta, nil
}

func internalServerError() error {
	return status.Errorf(codes.Internal, "internal server error.")
}

func uniqueOrInternalError(err error) error {
	if errors.Is(err, db.ErrUniqueViolation) {
		return status.Errorf(codes.AlreadyExists, "record already exists")
	}
	return internalServerError()
}

func notFoundOrInternalError(err error) error {
	if errors.Is(err, db.ErrRecordNotFound) {
		return status.Errorf(codes.NotFound, "record was not found")
	}
	return internalServerError()
}

func returnDBError(err error) error {
	if errors.Is(err, db.ErrRecordNotFound) {
		return status.Errorf(codes.NotFound, "record was not found")
	}
	if errors.Is(err, db.ErrUniqueViolation) {
		return status.Errorf(codes.AlreadyExists, "record already exists")
	}
	return internalServerError()
}
