package gapi

import (
	"context"
	"github.com/zvash/bgmood-circles-service/internal/cpb"
	"github.com/zvash/bgmood-circles-service/internal/db/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) React(ctx context.Context, req *cpb.ReactRequest) (*cpb.MoodResponse, error) {
	dto := cpbReactRequestToValReactRequest(req)
	if errs := server.validator.Validate(dto); errs != nil {
		return nil, errorResponsesToStatusErrors(errs)
	}
	userUUID, err := server.getOwnerUUID(ctx)
	if err != nil {
		return nil, err
	}
	moodUUID, err := stringToUUID(dto.MoodID)
	if err != nil {
		return nil, err
	}
	dbMood, err := server.db.GetMood(ctx, moodUUID)
	if err != nil {
		return nil, notFoundOrInternalError(err)
	}
	isMember, err := server.db.CheckIfMemberExists(ctx, repository.CheckIfMemberExistsParams{
		MemberID: userUUID,
		CircleID: dbMood.CircleID,
	})
	if err != nil {
		return nil, notFoundOrInternalError(err)
	}
	if !isMember {
		return nil, status.Errorf(codes.PermissionDenied, "unauthorized")
	}
	_, err = server.db.ReactToMood(ctx, repository.ReactToMoodParams{
		MoodID:     moodUUID,
		UserID:     userUUID,
		ReactionID: dto.ReactionID,
	})
	if err != nil {
		return nil, returnDBError(err)
	}
	mood, err := server.getMoodAndReactions(ctx, dbMood, userUUID)
	if err != nil {
		return nil, err
	}
	resp := &cpb.MoodResponse{
		Mood: mood,
	}
	return resp, nil
}
