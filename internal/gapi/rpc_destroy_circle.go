package gapi

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/zvash/bgmood-circles-service/internal/cpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) DestroyCircle(ctx context.Context, req *cpb.DestroyCircleRequest) (*cpb.DestroyCircleResponse, error) {
	idUUID, err := uuid.Parse(req.GetCircleId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not generate uuid from userId", err)
	}
	circle, err := server.getCircleForOwner(ctx, idUUID)
	if err != nil {
		return nil, err
	}
	err = server.db.RemoveCircleTransaction(ctx, idUUID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error.")
	}
	resp := &cpb.DestroyCircleResponse{
		Message: fmt.Sprintf("Successfully removed circle with id: %s", circle.ID.String()),
	}
	return resp, nil
}
