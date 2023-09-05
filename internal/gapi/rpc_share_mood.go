package gapi

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/zvash/bgmood-circles-service/internal/cpb"
	"github.com/zvash/bgmood-circles-service/internal/db"
	"github.com/zvash/bgmood-circles-service/internal/db/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

func (server *Server) CreateMood(ctx context.Context, req *cpb.CreateMoodRequest) (*cpb.MoodResponse, error) {
	dto := cpbCreateMoodRequestToValCreateMoodRequest(req)
	if errs := server.validator.Validate(dto); errs != nil {
		return nil, errorResponsesToStatusErrors(errs)
	}
	circleUUID, err := stringToUUID(dto.CircleID)
	if err != nil {
		return nil, err
	}
	userUUID, err := server.getOwnerUUID(ctx)
	if err != nil {
		return nil, err
	}

	ok, err := server.db.CheckIfMemberCanPostToCircle(ctx, repository.CheckIfMemberCanPostToCircleParams{
		MemberID: userUUID,
		CircleID: circleUUID,
	})
	if err != nil {
		return nil, notFoundOrInternalError(err)
	}
	if !ok {
		return nil, status.Errorf(codes.PermissionDenied, "unauthorized")
	}

	moodUUID, err := uuid.NewRandom()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not create a random uuid: %v", err)
	}
	mood, err := server.db.CreateMoodTransaction(ctx, db.CreateMoodTransactionParams{
		CreateMoodParams: repository.CreateMoodParams{
			ID:            moodUUID,
			CircleID:      circleUUID,
			PosterID:      userUUID,
			Image:         dto.Image,
			Description:   pgtype.Text{String: dto.Description, Valid: dto.Description != ""},
			SetBackground: dto.SetBackground,
		},
		AfterMoodCreated: func(mood repository.Mood) error {
			//Publish mood creation to circle members
			log.Printf("new mood was created: %v", mood)
			return nil
		},
	})
	if err != nil {
		return nil, returnDBError(err)
	}
	resp := &cpb.MoodResponse{
		Mood: &cpb.Mood{
			Id:            mood.ID.String(),
			PosterId:      mood.PosterID.String(),
			CircleId:      mood.CircleID.String(),
			Image:         mood.Image,
			Description:   &mood.Description.String,
			SetBackground: mood.SetBackground,
			CreatedAt:     timestamppb.New(mood.CreatedAt),
		},
	}
	return resp, nil
}
