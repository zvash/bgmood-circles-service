package gapi

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/zvash/bgmood-circles-service/internal/cpb"
	"github.com/zvash/bgmood-circles-service/internal/db/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateCircle(ctx context.Context, req *cpb.CreateCircleRequest) (*cpb.CreateCircleResponse, error) {
	dto := cpbCreateCircleRequestToValCreateCircleRequest(req)
	if errs := server.validator.Validate(dto); errs != nil {
		return nil, errorResponsesToStatusErrors(errs)
	}
	idUUID, err := uuid.NewRandom()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not generate random uuid", err)
	}
	owner, err := server.extractUserFromContext(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized.")
	}
	ownerUUID, err := uuid.Parse(owner.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not generate uuid from ownerId", err)
	}

	params := repository.CreateCircleParams{
		ID:          idUUID,
		OwnerID:     ownerUUID,
		Title:       dto.Title,
		Avatar:      pgtype.Text{String: dto.Avatar, Valid: true},
		Description: pgtype.Text{String: req.GetDescription(), Valid: req.GetDescription() != ""},
		CircleType:  repository.CircleType(req.GetCircleType()),
		IsPrivate:   req.GetIsPrivate(),
		IsFeatured:  false,
	}
	circle, err := server.db.CreateCircleTransaction(ctx, params)
	resp := &cpb.CreateCircleResponse{
		Circle: dbCircleToCPBCircle(circle),
	}
	return resp, nil
}
