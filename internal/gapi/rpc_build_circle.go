package gapi

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/zvash/bgmood-circles-service/internal/circlespb"
	"github.com/zvash/bgmood-circles-service/internal/db"
	"github.com/zvash/bgmood-circles-service/internal/db/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateCircle(ctx context.Context, req *circlespb.CreateCircleRequest) (*circlespb.CreateCircleResponse, error) {
	dto := circlespbCreateCircleRequestToValCreateCircleRequest(req)
	if errs := server.validator.Validate(dto); errs != nil {
		return nil, errorResponsesToStatusErrors(errs)
	}
	idUUID, err := uuid.NewRandom()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not generate random uuid", err)
	}
	ownerUUID, err := uuid.Parse(dto.OwnerId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not generate uuid from ownerId", err)
	}

	circle, err := server.db.CreateCircle(ctx, repository.CreateCircleParams{
		ID:          idUUID,
		OwnerID:     ownerUUID,
		Title:       dto.Title,
		Avatar:      pgtype.Text{String: dto.Avatar, Valid: true},
		Description: pgtype.Text{String: req.GetDescription(), Valid: req.GetDescription() != ""},
		CircleType:  repository.CircleType(req.GetCircleType()),
		IsPrivate:   req.GetIsPrivate(),
		IsFeatured:  false,
	})
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			return nil, status.Errorf(codes.AlreadyExists, "a circle with the same name already exists")
		}
		return nil, status.Errorf(codes.Internal, "internal server error.")
	}
	resp := &circlespb.CreateCircleResponse{
		Circle: dbCircleToCriclespbCircle(circle),
	}
	return resp, nil
}
