package gapi

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/zvash/bgmood-circles-service/internal/db"
	"github.com/zvash/bgmood-circles-service/internal/db/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
