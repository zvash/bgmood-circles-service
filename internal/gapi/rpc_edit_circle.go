package gapi

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/zvash/bgmood-circles-service/internal/cpb"
	"github.com/zvash/bgmood-circles-service/internal/db"
	"github.com/zvash/bgmood-circles-service/internal/db/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) EditCircle(ctx context.Context, req *cpb.EditCircleRequest) (*cpb.EditCircleResponse, error) {
	dto := cpbEditCircleRequestToValEditCircleRequest(req)
	if errs := server.validator.Validate(dto); errs != nil {
		return nil, errorResponsesToStatusErrors(errs)
	}
	dbCircle, err := server.db.UpdateCircle(ctx, repository.UpdateCircleParams{
		Title:       pgtype.Text{String: dto.Title, Valid: req.Title != nil},
		Avatar:      pgtype.Text{String: dto.Avatar, Valid: req.Avatar != nil},
		Description: pgtype.Text{String: dto.Description, Valid: req.Description != nil},
		IsPrivate:   pgtype.Bool{Bool: dto.IsPrivate, Valid: req.IsPrivate != nil},
		CircleType: repository.NullCircleType{
			CircleType: repository.CircleType(dto.CircleType),
			Valid:      req.CircleType != nil,
		},
	})
	if err != nil {
		if errors.Is(err, db.ErrUniqueViolation) {
			return nil, status.Errorf(codes.AlreadyExists, "a circle with the provided title already exists")
		}
		return nil, status.Errorf(codes.Internal, "internal server error.")
	}
	resp := &cpb.EditCircleResponse{
		Circle: dbCircleToCPBCircle(dbCircle),
	}
	return resp, nil
}
