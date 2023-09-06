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
	circle, err := server.db.DisplayCircleForUser(ctx, repository.DisplayCircleForUserParams{
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
		ViewerIsMember: circle.IsMember.Bool,
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

func (server *Server) getExploredCirclesMeta(ctx context.Context, userID uuid.UUID, page int64) ([]*cpb.CircleMeta, error) {
	var limit int32 = 10
	circles, err := server.db.ExploreCirclesForUserPaginated(ctx, repository.ExploreCirclesForUserPaginatedParams{
		MemberID: userID,
		Offset:   int32(page-1) * limit,
		Limit:    limit,
	})
	if err != nil {
		return nil, notFoundOrInternalError(err)
	}
	circlesMeta := make([]*cpb.CircleMeta, 0)
	for _, circle := range circles {
		canPost, err := server.db.CheckIfMemberCanPostToCircle(ctx, repository.CheckIfMemberCanPostToCircleParams{
			MemberID: userID,
			CircleID: circle.ID,
		})
		if err != nil {
			continue
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
			MoodCount:      circle.MoodCount,
			MemberCount:    circle.MemberCount,
			ViewerIsMember: circle.IsMember,
			ViewerIsOwner:  circle.OwnerID == userID,
			ViewerCanPost:  canPost,
		}
		circlesMeta = append(circlesMeta, circleMeta)
	}
	return circlesMeta, nil
}

func (server *Server) prepareChangingAccessToCircle(ctx context.Context, circleID, userID string) (circleUUID uuid.UUID, memberUUID uuid.UUID, err error) {
	circleUUID, err = stringToUUID(circleID)
	if err != nil {
		return
	}
	userUUID, err := server.getOwnerUUID(ctx)
	if err != nil {
		return
	}
	ok, err := server.db.CheckIfMemberCanChangeUsersAccess(ctx, repository.CheckIfMemberCanChangeUsersAccessParams{
		MemberID: userUUID,
		CircleID: circleUUID,
	})
	if err != nil {
		err = notFoundOrInternalError(err)
		return
	}
	if !ok {
		err = status.Errorf(codes.PermissionDenied, "unauthorized")
		return
	}
	memberUUID, err = stringToUUID(userID)
	if err != nil {
		return
	}
	return
}

func (server *Server) getMoodAndReactions(ctx context.Context, mood repository.Mood, userID uuid.UUID) (*cpb.Mood, error) {
	moodID := mood.ID
	moodReactions, err := server.db.GetMoodReactions(ctx, moodID)
	if err != nil {
		return nil, notFoundOrInternalError(err)
	}
	var description *string
	if mood.Description.Valid {
		description = &mood.Description.String
	}
	reactions := make([]*cpb.ReactionAggregated, 0)
	for _, moodReaction := range moodReactions {
		var emoji *string
		var textRepresentation *string
		if moodReaction.Emoji.Valid {
			emoji = &moodReaction.Emoji.String
		}
		if moodReaction.TextRepresentation.Valid {
			textRepresentation = &moodReaction.TextRepresentation.String
		}
		r := &cpb.ReactionAggregated{
			Id:                 moodReaction.ID,
			Count:              moodReaction.ReactionCount,
			Name:               moodReaction.Name,
			Emoji:              emoji,
			TextRepresentation: textRepresentation,
		}
		reactions = append(reactions, r)
	}
	var userMoodReaction *cpb.Reaction
	dbUserMoodReaction, err := server.db.GetUserMoodReaction(ctx, repository.GetUserMoodReactionParams{
		MoodID: moodID,
		UserID: userID,
	})
	if err == nil {
		var emoji *string
		var textRepresentation *string
		if dbUserMoodReaction.Emoji.Valid {
			emoji = &dbUserMoodReaction.Emoji.String
		}
		if dbUserMoodReaction.TextRepresentation.Valid {
			textRepresentation = &dbUserMoodReaction.TextRepresentation.String
		}
		userMoodReaction = &cpb.Reaction{
			Id:                 dbUserMoodReaction.ID,
			Name:               dbUserMoodReaction.Name,
			Emoji:              emoji,
			TextRepresentation: textRepresentation,
		}
	}
	res := &cpb.Mood{
		Id:              mood.ID.String(),
		PosterId:        mood.PosterID.String(),
		CircleId:        mood.CircleID.String(),
		Image:           mood.Image,
		SetBackground:   mood.SetBackground,
		CreatedAt:       timestamppb.New(mood.CreatedAt),
		Description:     description,
		UserReaction:    userMoodReaction,
		OthersReactions: reactions,
	}
	return res, nil
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
