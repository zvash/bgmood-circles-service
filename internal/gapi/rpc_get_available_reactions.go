package gapi

import (
	"context"
	"github.com/zvash/bgmood-circles-service/internal/cpb"
)

func (server *Server) GetAvailableReactions(ctx context.Context,
	req *cpb.GetAvailableReactionsRequest,
) (*cpb.GetAvailableReactionsResponse, error) {
	reactions := make([]*cpb.Reaction, 0)
	dbReactions, err := server.db.GetAvailableReactions(ctx)
	if err != nil {
		return nil, notFoundOrInternalError(err)
	}
	for _, dbReaction := range dbReactions {
		var emoji *string
		var textRepresentation *string
		if dbReaction.Emoji.Valid {
			emoji = &dbReaction.Emoji.String
		}
		if dbReaction.TextRepresentation.Valid {
			textRepresentation = &dbReaction.TextRepresentation.String
		}
		reaction := &cpb.Reaction{
			Id:                 dbReaction.ID,
			Name:               dbReaction.Name,
			Emoji:              emoji,
			TextRepresentation: textRepresentation,
		}
		reactions = append(reactions, reaction)
	}
	resp := &cpb.GetAvailableReactionsResponse{
		Reactions: reactions,
	}
	return resp, nil
}
