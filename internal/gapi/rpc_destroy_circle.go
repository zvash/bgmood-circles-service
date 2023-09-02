package gapi

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/zvash/bgmood-circles-service/internal/circlespb"
	"github.com/zvash/bgmood-circles-service/internal/db"
	"github.com/zvash/bgmood-circles-service/internal/db/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) DestroyCircle(ctx context.Context, req *circlespb.DestroyCircleRequest) (*circlespb.DestroyCircleResponse, error) {
	user, err := server.extractUserFromContext(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized.")
	}
	idUUID, err := uuid.Parse(req.GetCircleId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not generate uuid from userId", err)
	}
	userUUID, err := uuid.Parse(user.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not generate uuid from circleId", err)
	}
	circle, err := server.db.GetUserCircle(ctx, repository.GetUserCircleParams{
		ID:      idUUID,
		OwnerID: userUUID,
	})
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "could not find the circle")
		}
		return nil, status.Errorf(codes.Internal, "internal server error.")
	}
	err = server.db.RemoveCircleTransaction(ctx, idUUID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error.")
	}
	resp := &circlespb.DestroyCircleResponse{
		Message: fmt.Sprintf("Successfully removed circle with id: %s", circle.ID.String()),
	}
	return resp, nil
}
